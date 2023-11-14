pub mod user_gen {
    tonic::include_proto!("user");
}

pub mod general {
    tonic::include_proto!("general");
}


pub use user_gen::user_rpc_server::{UserRpcServer, UserRpc};
use user_gen::{CreateUserPayload, UpdateUserPayload, GetUserPayload, GetUserResponse};
use general::IsSuccess;
use sea_orm::DatabaseConnection;
use tonic::{Request,  Response, Status};
use logger::info;
use crate::repository::UserInfoRepository;

#[derive(Debug)]
pub struct UserService {
    user_info_repository: UserInfoRepository
}

impl UserService {
    pub fn new(db: DatabaseConnection) -> Self {
        return Self {
            user_info_repository: UserInfoRepository::new(db)
        }
    }
}

#[tonic::async_trait]
impl UserRpc for UserService {
    async fn create_user(&self, request: Request<CreateUserPayload>) ->  Result<Response<IsSuccess>, Status> {
        info!("Got request! {:?}", request);

        let req = request.into_inner();

        let res = self.user_info_repository.create(req.user_id, req.picture).await;
        match res {
            Ok(is_success) => {
                info!("Successfully created user!");
                let res = IsSuccess {
                    is_success
                };
                Ok(Response::new(res))
            },
            Err(_) => {
                let status = Status::new(tonic::Code::Internal, "internal error");
                Err(status)
            }
        }
    }

    async fn update_user(&self, request: Request<UpdateUserPayload>) -> Result<Response<IsSuccess>, Status> {
        Ok(Response::new(IsSuccess { is_success: true }))
    }

    async fn get_user(&self, request: Request<GetUserPayload>) -> Result<Response<GetUserResponse>, Status> {
        Ok(Response::new(GetUserResponse::default()))
    }
}
