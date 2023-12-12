pub mod wasm;
pub mod nats;
pub mod grpc;


pub type Result<T> = std::result::Result<T, Box<dyn std::error::Error>>;
