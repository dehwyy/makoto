mod service;
mod repository;
mod utils;

use tonic::transport::Server;
use service::service::AuthRpcServiceImplementation;

use makoto_grpc::pkg::api_auth::auth_rpc_server::AuthRpcServer;
use makoto_logger::{Logger, info};
use makoto_lib::Result as MakotoResult;

#[tokio::main]
async fn main() -> MakotoResult<()> {

    Logger::init()?;

    let hosts = makoto_config::hosts::Hosts::new();
    let addr = hosts.auth.parse()?;

    let db = makoto_db::new().await.expect("cannot open database connection");

    let credentials_repo = repository::credentials::Credentials::new(db.clone());
    let tokens_repo = repository::token::Tokens::new(db.clone());
    let oauth_repo = repository::oauth::Oauth::new(db.clone());

    let auth_service= AuthRpcServiceImplementation::new(credentials_repo, tokens_repo, oauth_repo);
    let auth_service = AuthRpcServer::new(auth_service);

    info!("server start! host: {}", addr);

    // settings for grpc-web library
    Server::builder()
        .add_service(auth_service)
        .serve(addr)
        .await?;

    info!("server stoped!");

    Ok(())
}
