use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(tag = "type")]
pub enum Packet {
    #[serde(rename = "login")]
    Login { token: String },
    
    #[serde(rename = "move")]
    Move { x: f32, y: f32, z: f32 },
    
    #[serde(rename = "heartbeat")]
    Heartbeat,
    
    #[serde(rename = "world_update")]
    WorldUpdate { players: Vec<PlayerState> },
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlayerState {
    pub id: u32,
    pub x: f32,
    pub y: f32,
    pub z: f32,
}
