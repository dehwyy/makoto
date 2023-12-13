mod service;

use tonic::transport::Server;
use service::service::ApiRpcServiceImplementation;

use makoto_grpc::pkg::api::api_rpc_server::ApiRpcServer;
use makoto_logger::{Logger, info};
use makoto_lib::Result as MakotoResult;

#[tokio::main]
async fn main() -> MakotoResult<()> {

    Logger::init()?;

    let hosts = makoto_config::hosts::Hosts::new();
    let addr = hosts.gateway.parse()?;

    let api = ApiRpcServiceImplementation::new();
    let api_service = ApiRpcServer::new(api);

    info!("server start! host: {}", addr);

    // settings for grpc-web library
    Server::builder()
        .accept_http1(true)
        .layer(tonic_web::GrpcWebLayer::new())
        .add_service(tonic_web::enable(api_service))
        .serve(addr)
        .await?;

    info!("server stoped!");

    Ok(())
}
