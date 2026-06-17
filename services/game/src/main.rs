mod net;
mod world;
mod systems;

use std::sync::Arc;
use tokio::net::UdpSocket;
use std::time::Duration;
use crate::world::state::{WorldState, SharedWorldState};
use crate::net::packet::{Packet, PlayerState};
use dotenvy::dotenv;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenv().ok(); 
    
    let state: SharedWorldState = Arc::new(WorldState::new());
    
    let udp_addr = std::env::var("UDP_ADDR").unwrap_or_else(|_| "0.0.0.0:8081".to_string());
    let state_for_udp = state.clone();
    
    tokio::spawn(async move {
        if let Err(e) = net::udp::run_udp_server(&udp_addr, state_for_udp).await {
            eprintln!("UDP Server error: {}", e);
        }
    });
    
    let state_for_timeout = state.clone();
    tokio::spawn(async move {
        systems::timeout::run_timeout_system(state_for_timeout).await;
    });
    
    let socket = UdpSocket::bind("0.0.0.0:0").await?; 
    let state_for_broadcast = state.clone();
    
    println!("Broadcast loop started");
    
    loop {
        tokio::time::sleep(Duration::from_millis(50)).await;
        
        let mut player_states = Vec::new();
        for entry in state_for_broadcast.players.iter() {
            let p = entry.value();
            player_states.push(PlayerState {
                id: p.id,
                x: p.x,
                y: p.y,
                z: p.z,
            });
        }
        
        if player_states.is_empty() {
            continue;
        }
        
        let packet = Packet::WorldUpdate { players: player_states };
        let serialized = serde_json::to_vec(&packet)?;
        
        for entry in state_for_broadcast.players.iter() {
            let addr = entry.key();
            let _ = socket.send_to(&serialized, addr).await;
        }
    }
}