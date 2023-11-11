pub mod user {
    tonic::include_proto!("user");
}

pub mod general {
    tonic::include_proto!("general");
}

use tonic::{transport::Server,  Request,  Response, Status};

use user::{user_rpc_server::{UserRpcServer, UserRpc}};
use user::{CreateUserPayload};
use general::IsSuccess;

#[derive(Debug, Default)]
pub struct UserService {}

#[tonic::async_trait]
impl UserRpc for UserService {
    async fn create_user(&self, request: Request<CreateUserPayload>) ->  Result<Response<IsSuccess>, Status> {
        println!("requset");

        Ok(Response::new(IsSuccess { is_success: true }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:4010".parse()?;
    let user_service = UserService::default();

    Server::builder()
        .add_service(UserRpcServer::new(user_service))
        .serve(addr)
        .await?;

    Ok(())
}
