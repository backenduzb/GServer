use tokio::sync::mpsc;

pub struct Client {
	pub id: u32,
	pub sender: mpsc::UnboundedSender<String>,
}