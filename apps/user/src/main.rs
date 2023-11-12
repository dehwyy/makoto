pub mod user {
    tonic::include_proto!("user");
}

pub mod general {
    tonic::include_proto!("general");
}

use std::net::ToSocketAddrs;

use tonic::{transport::Server,  Request,  Response, Status};
use user::user_rpc_server::{UserRpcServer, UserRpc};
use user::{CreateUserPayload};
use general::IsSuccess;

use database::Database;
use config::Config;

#[derive(Debug, Default)]
pub struct UserService {}

#[tonic::async_trait]
impl UserRpc for UserService {
    async fn create_user(&self, request: Request<CreateUserPayload>) ->  Result<Response<IsSuccess>, Status> {
        println!("requset {:?}", request);

        Ok(Response::new(IsSuccess { is_success: true }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let config = Config::new();


    // let db = Database::new(String::from("")).await?;

    let addr = config.user_url.split("//").collect::<Vec<_>>()[1].to_socket_addrs()?.next().unwrap();



    let user_service = UserService::default();

        Server::builder()
        .add_service(UserRpcServer::new(user_service))
        .serve(addr)
        .await?;

    Ok(())
}
