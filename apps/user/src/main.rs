mod repository;
mod grpc;

pub mod result {
    pub type T<R, E = Box<dyn std::error::Error>> = Result<R, E>;
}

use std::net::ToSocketAddrs;
use tonic::transport::Server;
use config::Config;
use database::Database;
use grpc::{UserService, UserRpcServer};
use logger::{Logger, info};

#[tokio::main]
async fn main() -> result::T<()> {
    let config = Config::new();
    Logger::init().expect("Cannot run logger!");

    let db = Database::new(config.database_url).await?;

    // addr looks like http(s)://localhost:4000 -> only localhost:4000 needed
    let addr = config.user_url.split("//").collect::<Vec<_>>()[1].to_socket_addrs()?.next().unwrap();

    info!("Server started on port {}", addr);
    let user_service = UserService::new(db);
        Server::builder()
        .add_service(UserRpcServer::new(user_service))
        .serve(addr)
        .await?;

    Ok(())
}
