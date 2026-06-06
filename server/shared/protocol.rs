use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
#[serde(tag="t")]
pub enum ClientPacket {
	Join {
		name: String,
	},

	Move {
		x: f32,
		y: f32,
		z: f32,
	},
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(tag="t")]
pub enum ServerPacket {
	Welcome {
		id: u32,
	},

	PlayerJoined {
		id: u32,
		name: String,
	},

	PlayerMoved {
		id: u32,
		x: f32,
		y: f32,
		z: f32,
	},

	PlayerLeft {
		id: u32,
	},
}