use std::collections::HashMap;

use super::player::Player;

pub struct World {
	pub players: HashMap<u32, Player>,
}

impl World {
	pub fn new() -> Self {
		Self {
			players: HashMap::new(),
		}
	}
}