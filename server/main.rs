mod net {
	pub mod client;
	pub mod listener;
	pub mod packets;
}

mod systems {
	pub mod movement;
}

mod world {
	pub mod player;
	pub mod world;
}

mod shared {
	pub mod protocol;
}

use std::{collections::HashMap, sync::Arc};

use net::listener::{run, Clients};
use tokio::sync::RwLock;
use world::world::World;

#[tokio::main]
async fn main() {
	let world = Arc::new(RwLock::new(World::new()));
	let clients: Clients = Arc::new(RwLock::new(HashMap::new()));

	run(world, clients).await;
}