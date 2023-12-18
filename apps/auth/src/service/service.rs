use makoto_grpc::pkg::api_auth::auth_rpc_server::AuthRpc;
use makoto_grpc::pkg::api_auth::{
  SignInRequest, AuthorizationResponse, SignUpRequest,
  SignInOauthRequest, SignInTokenRequest, SignOutRequest,
  ConfirmMailRequest, SubmitNewPasswordByCodeRequest,
  IsEmailAvailableRequest, IsUsernameAvailableRequest
};

use makoto_grpc::Result as GrpcResult;
use makoto_grpc::pkg::general::BoolStatus;
use makoto_logger::{info, warn};

use sea_orm::prelude::Uuid;
use tokio::join;
use tonic::{Request as Req, Status, Response, async_trait};

use crate::repository::credentials::{Credentials, UserPayload};
use crate::repository::token::Tokens;
use crate::repository::oauth::Oauth;

use crate::utils::hasher::Hasher;
use crate::utils::validator::Validator;

pub struct AuthRpcServiceImplementation {
  pub credentials_repository: Credentials,
  pub tokens_repository: Tokens,
  pub oauth_repository: Oauth
}

impl AuthRpcServiceImplementation {
  pub fn new(credentials_repo: Credentials, tokens_repo: Tokens, oauth_repo: Oauth) -> Self {
    Self {
      credentials_repository: credentials_repo,
      tokens_repository: tokens_repo,
      oauth_repository: oauth_repo
    }
  }
}

#[async_trait]
impl AuthRpc for AuthRpcServiceImplementation {
  async fn sign_up(&self, req: Req<SignUpRequest>) -> GrpcResult<AuthorizationResponse> {
    let req = req.into_inner();

    // validation
    {
      Validator::username(&req.username).map_err(|msg| Status::invalid_argument(msg))?;
      Validator::email(&req.email).map_err(|msg| Status::invalid_argument(msg))?;
      Validator::password(&req.password).map_err(|msg| Status::invalid_argument(msg))?;
    }

    // check availability
    {
      let (username, email) =  join!(
        self.credentials_repository.is_username_available(&req.username),
        self.credentials_repository.is_email_available(&req.email)
      );

      let is_available = username.map_err(|msg| Status::internal(msg))?;
      if !is_available {
        return Err(Status::already_exists("username is already taken"));
      }

      let is_available = email.map_err(|msg| Status::internal(msg))?;
      if !is_available {
        return Err(Status::already_exists("email is already taken"));
      }
    }

    // new user_id
    let user_id = Uuid::new_v4();

    // create user
    self.credentials_repository.create_user(UserPayload {
      user_id: user_id.clone(),
      username: req.username.clone(),
      email: req.email,
      password: Hasher::hash(req.password).expect("cannot hash password"),
    }).await.map_err(|msg| Status::internal(msg))?;

    // generate tokens
    let new_access_token = self.tokens_repository.create_new_token_pair(user_id.clone(), &req.username).await.map_err(|msg| Status::internal(msg))?;

    // initialize empty oauth
    // self.oauth_repository.create_empty_record(user_id.clone()).await.map_err(|msg| Status::internal(msg))?;

    Ok(Response::new(
      AuthorizationResponse {
        token: new_access_token,
        username: req.username,
        used_id: user_id.to_string()
      }
    ))
  }

  async fn sign_in(&self, req: Req<SignInRequest>) -> GrpcResult<AuthorizationResponse> {
    let req = req.into_inner();

    let user = match req.username.is_empty() {
      true => self.credentials_repository.get_user_by_email(&req.email).await,
      false => self.credentials_repository.get_user_by_username(&req.username).await
    }.map_err(|err| Status::not_found(err))?;

    let password = match user.password {
      Some(password) => password,
      None => return Err(Status::not_found("empty password (trying to signin using oauth2 user)"))
    };

    // check password
    if !Hasher::verify(&req.password, &password).map_err(|_| Status::internal("cannot verify password (hasher error)"))? {
      warn!("password is incorrect for user {}", user.username);
      return Err(Status::unauthenticated("password is incorrect"));
    }

    // generate new access_token
    let new_access_token = self.tokens_repository.create_new_access_token(user.id.clone(), &user.username).await.map_err(|msg| Status::internal(msg))?;

    Ok(Response::new(
      AuthorizationResponse {
        token: new_access_token,
        username: user.username,
        used_id: user.id.to_string()
      }
    ))
  }

  async fn sign_in_oauth(&self, req: Req<SignInOauthRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn sign_in_token(&self, req: Req<SignInTokenRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn sign_out(&self, req: Req<SignOutRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }

  async fn confirm_email(&self, req: Req<ConfirmMailRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }

  async fn update_password(&self, req: Req<SubmitNewPasswordByCodeRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn is_email_available(&self, req: Req<IsEmailAvailableRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }

  async fn is_username_available(&self, req: Req<IsUsernameAvailableRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }
}

#[cfg(test)]
mod tests {
  use std::env;

use super::*;
  use lazy_static::lazy_static;

  async fn get_test_service() -> AuthRpcServiceImplementation {
    let db = makoto_db::get_test_db().await;
    let credentials = Credentials::new(db.clone());
    let tokens = Tokens::new(db.clone());
    let oauth = Oauth::new(db.clone());

    AuthRpcServiceImplementation::new(credentials, tokens, oauth)
  }


  #[tokio::test]
  async fn test_signup_signin_flow() {
    let service = get_test_service().await;

    service.sign_up(Req::new(SignUpRequest {
      username: "dehwyy".to_string(),
      email: "dehwyy@qqq.com".to_string(),
      password: "some_pass".to_string()
    }))
      .await.map_err(|err| {
        eprintln!("{err:?}")
      })
      .unwrap();

      service.sign_in(Req::new(SignInRequest {
        username: "dehwyy".to_string(),
        email: "dehwyy@qqq.com".to_string(),
        password: "some_pass".to_string()
      }))
        .await.map_err(|err| {
          eprintln!("{err:?}")
        })
        .unwrap();
  }
}
