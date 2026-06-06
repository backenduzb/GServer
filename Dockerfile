FROM rust:1.78 as builder

WORKDIR /service

COPY Cargo.toml ./
COPY server ./server

RUN cargo build --relase --bin server

FROM debian:bookworm-slim

WORKDIR /service

RUN apt-get update \ 
	&& apt-get install -y ca-certificates \
	&& rm -rf /var/lib/apt/libraries/lists/*
COPY --from=builder /service/target/relase/server .

EXPOSE 5555

CMD ["./server"]