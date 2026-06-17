FROM rust:1.85-slim AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y pkg-config libssl-dev && rm -rf /var/lib/apt/lists/*

COPY . .
RUN cargo build --release

FROM debian:bookworm-slim

WORKDIR /app
RUN apt-get update && apt-get install -y libssl3 ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/target/release/main /app/game-server
COPY .env /app/.env

ENV RUST_LOG=info
EXPOSE 8081/udp

CMD ["./game-server"]