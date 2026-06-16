FROM rust:1.75 AS builder

WORKDIR /app

COPY . .
RUN cargo build --release

FROM gcr.io/distroless/cc-debian12

WORKDIR /app

RUN apt-get update && apt-get install -y\
	ca-certificates\
	&& rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/target/release/microservice /app/microservice

ENV RUST_LOG=info
EXPOSE 7777/udp 

CMD ["./microservice"]