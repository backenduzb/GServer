use std::net::SocketAddr;
use tokio::net::UdpSocket;
use std::sync::Arc;
use crate::net::packet::{Packet, PlayerState};
use crate::world::state::SharedWorldState;
use crate::world::player::Player;
use jsonwebtoken::{decode, DecodingKey, Validation, Algorithm};
use serde::{Deserialize, Serialize};
use chrono::Utc;

#[derive(Debug, Serialize, Deserialize)]
struct Claims {
    user_id: u32,
    exp: usize,
}

pub async fn run_udp_server(addr: &str, state: SharedWorldState) -> anyhow::Result<()> {
    let socket = Arc::new(UdpSocket::bind(addr).await?);
    println!("UDP Server listening on {} ", addr);

    let mut buf = [0u8; 1024];

    loop {
        let (len, addr) = socket.recv_from(&mut buf).await?;
        let data = &buf[..len];

        let state_clone = state.clone();

        if let Ok(packet) = serde_json::from_slice::<Packet>(data) {
            match packet {
                Packet::Login { token } => {
                    let secret = std::env::var("JWT_SECRET").unwrap_or_else(|_| "secret".to_string());
                    let validation = Validation::new(Algorithm::HS256);
                    
                    if let Ok(token_data) = decode::<Claims>(
                        &token,
                        &DecodingKey::from_secret(secret.as_bytes()),
                        &validation,
                    ) {
                        let user_id = token_data.claims.user_id;
                        state_clone.players.insert(addr, Player::new(user_id, addr));
                        println!("Player {} logged in from {} ", user_id, addr);
                    }
                }
                Packet::Move { x, y, z } => {
                    if let Some(mut player) = state_clone.players.get_mut(&addr) {
                        player.x = x;
                        player.y = y;
                        player.z = z;
                        player.last_seen = Utc::now();
                    }
                }
                Packet::Heartbeat => {
                    if let Some(mut player) = state_clone.players.get_mut(&addr) {
                        player.last_seen = Utc::now();
                    }
                }
                _ => {}
            }
        }
    }
}