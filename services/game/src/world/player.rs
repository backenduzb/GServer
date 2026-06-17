use std::net::SocketAddr;
use chrono::{DateTime, Utc};

#[derive(Debug, Clone)]
pub struct Player {
    pub id: u32,
    pub addr: SocketAddr,
    pub x: f32,
    pub y: f32,
    pub z: f32,
    pub last_seen: DateTime<Utc>,
}

impl Player {
    pub fn new(id: u32, addr: SocketAddr) -> Self {
        Self {
            id,
            addr,
            x: 0.0,
            y: 0.0,
            z: 0.0,
            last_seen: Utc::now(),
        }
    }
}
