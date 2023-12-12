mod service;

use tonic::transport::Server;
use service::service::ApiRpcServiceImplementation;
use grpc::pkg::api::api_rpc_server::ApiRpcServer;

#[tokio::main]
async fn main() {
    let addr = "127.0.0.1:8080".parse().unwrap();

    let api = ApiRpcServiceImplementation::new();
    let api_service = ApiRpcServer::new(api);


    println!("server listen on: {}", addr);

    Server::builder()
        .accept_http1(true)
        .layer(tonic_web::GrpcWebLayer::new())
        .add_service(tonic_web::enable(api_service))
        .serve(addr)
        .await.unwrap();

    println!("server shutdown!");
}
