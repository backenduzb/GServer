use std::time::Duration;
use tokio::time::sleep;
use crate::world::state::SharedWorldState;
use chrono::Utc;
use reqwest::Client;
use serde::Serialize;

#[derive(Serialize)]
struct PositionUpdate {
    user_id: u32,
    x: f32,
    y: f32,
    z: f32,
}

pub async fn run_timeout_system(state: SharedWorldState) {
    let client = Client::new();
    let go_app_url = std::env::var("GO_APP_URL").unwrap_or_else(|_| "http://localhost:8080".to_string());
    
    loop {
        sleep(Duration::from_secs(1)).await;
        
        let now = Utc::now();
        let mut to_remove = Vec::new();
        
        {
            for entry in state.players.iter() {
                let player = entry.value();
                if now.signed_duration_since(player.last_seen).num_seconds() > 5 {
                    to_remove.push((entry.key().clone(), player.clone()));
                }
            }
        }
        
        for (addr, player) in to_remove {
            println!("Player {} timed out from {}", player.id, addr);
            state.players.remove(&addr);
            
            let update = PositionUpdate {
                user_id: player.id,
                x: player.x,
                y: player.y,
                z: player.z,
            };
            
            let url = format!("{}/pos/internal/update", go_app_url);
            match client.put(&url)
                .json(&update)
                .send()
                .await {
                Ok(res) => {
                    if !res.status().is_success() {
                        eprintln!("Failed to sync player {} to Go: {}", player.id, res.status());
                    }
                }
                Err(e) => eprintln!("Error sending timeout sync to Go: {}", e),
            }
        }
    }
}