use crate::shared::protocol::ServerPacket;

pub fn serialize(packet: &ServerPacket) -> String {
	format!(
		"{}\n",
		serde_json::to_string(packet).unwrap()
	)
}