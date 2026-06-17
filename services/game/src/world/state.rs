use dashmap::DashMap;
use std::net::SocketAddr;
use std::sync::Arc;
use crate::world::player::Player;

pub struct WorldState {
    pub players: DashMap<SocketAddr, Player>,
}

impl WorldState {
    pub fn new() -> Self {
        Self {
            players: DashMap::new(),
        }
    }
}

pub type SharedWorldState = Arc<WorldState>;
