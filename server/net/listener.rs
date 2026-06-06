use std::{
    collections::HashMap,
    sync::{
        atomic::{AtomicU32, Ordering},
        Arc,
    },
};

use tokio::{
    io::{AsyncBufReadExt, AsyncWriteExt, BufReader},
    net::{TcpListener, TcpStream},
    sync::{mpsc, RwLock},
};

use crate::{
    shared::protocol::{ClientPacket, ServerPacket},
    world::{
        player::Player,
        world::World,
    },
};

static NEXT_ID: AtomicU32 = AtomicU32::new(1);

pub type Clients =
    Arc<RwLock<HashMap<u32, mpsc::UnboundedSender<String>>>>;

pub async fn run(
    world: Arc<RwLock<World>>,
    clients: Clients,
) {
    let listener =
        TcpListener::bind("0.0.0.0:5555")
            .await
            .unwrap();

    println!("Server listening on :5555");

    loop {
        let (socket, _) =
            listener.accept().await.unwrap();

        let world = world.clone();
        let clients = clients.clone();

        tokio::spawn(async move {
            handle(socket, world, clients).await;
        });
    }
}

async fn handle(
    socket: TcpStream,
    world: Arc<RwLock<World>>,
    clients: Clients,
) {
    let id = NEXT_ID.fetch_add(
        1,
        Ordering::Relaxed,
    );

    let (reader, mut writer) =
        socket.into_split();

    let (tx, mut rx) =
        mpsc::unbounded_channel::<String>();

    clients.write().await.insert(id, tx);

    tokio::spawn(async move {
        while let Some(msg) = rx.recv().await {
            let _ =
                writer.write_all(msg.as_bytes()).await;
        }
    });

    let mut lines =
        BufReader::new(reader).lines();

    while let Ok(Some(line)) =
        lines.next_line().await
    {
        let Ok(packet) =
            serde_json::from_str::<ClientPacket>(&line)
        else {
            continue;
        };

        match packet {
            ClientPacket::Join { name } => {
                world.write().await.players.insert(
                    id,
                    Player {
                        id,
                        name: name.clone(),
                        x: 0.0,
                        y: 0.0,
                        z: 0.0,
                    },
                );

                broadcast(
                    &clients,
                    ServerPacket::PlayerJoined {
                        id,
                        name,
                    },
                )
                .await;
            }

            ClientPacket::Move { x, y, z } => {
                if let Some(player) =
                    world.write().await.players.get_mut(&id)
                {
                    player.x = x;
                    player.y = y;
                    player.z = z;
                }

                broadcast(
                    &clients,
                    ServerPacket::PlayerMoved {
                        id,
                        x,
                        y,
                        z,
                    },
                )
                .await;
            }
        }
    }

    world.write().await.players.remove(&id);

    clients.write().await.remove(&id);

    broadcast(
        &clients,
        ServerPacket::PlayerLeft { id },
    )
    .await;
}

async fn broadcast(
    clients: &Clients,
    packet: ServerPacket,
) {
    let msg = format!(
        "{}\n",
        serde_json::to_string(&packet).unwrap()
    );

    for tx in clients.read().await.values() {
        let _ = tx.send(msg.clone());
    }
}