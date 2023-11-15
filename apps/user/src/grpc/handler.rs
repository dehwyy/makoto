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
use crate::repository::{UserInfoRepository, FindUserResult, UpdateUserModel};

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
        let req = request.into_inner();

        let payload = UpdateUserModel {
            dark_bg: req.dark_bg,
            light_bg: req.light_bg,
            picture: req.picture,
            description: req.description,
            languages: req.languages
        };

        let res = self.user_info_repository.update(&req.user_id, payload).await;
        match res {
            Ok(is_success) => {
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

    async fn get_user(&self, request: Request<GetUserPayload>) -> Result<Response<GetUserResponse>, Status> {
        let req = request.into_inner();

        let res = self.user_info_repository.get(req.user_id.clone()).await.unwrap_or(FindUserResult::NotFound);

        match res {
            FindUserResult::Found(user) => {
                let res = GetUserResponse {
                    dark_bg: user.user.background_dark.unwrap_or_default(),
                    light_bg: user.user.background_light.unwrap_or_default(),
                    description: user.user.description.unwrap_or_default(),
                    picture: user.user.picture.unwrap_or_default(),
                    languages: user.languages.iter().map(|language| language.lang.clone())
                                                    .collect::<Vec<_>>()
                };

                Ok(Response::new(res))
            },
            FindUserResult::NotFound => {
                let status = Status::new(tonic::Code::NotFound, format!("User with id {} wasn't found", req.user_id));
                Err(status)
            }
        }
    }
}
