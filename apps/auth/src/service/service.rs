use makoto_grpc::pkg::api_auth::auth_rpc_server::AuthRpc;
use makoto_grpc::pkg::api_auth::{
  SignInRequest, AuthorizationResponse, SignUpRequest,
  SignInOauthRequest, SignInTokenRequest, SignOutRequest,
  ConfirmMailRequest, SubmitNewPasswordByCodeRequest,
  IsEmailAvailableRequest, IsUsernameAvailableRequest
};

use makoto_grpc::Result as GrpcResult;
use makoto_grpc::pkg::general::BoolStatus;
use tonic::Request as Req;

#[derive(Default)]
pub struct AuthRpcServiceImplementation;

impl AuthRpcServiceImplementation {
  pub fn new() -> Self {
    Self {
    }
  }
}

#[tonic::async_trait]
impl AuthRpc for AuthRpcServiceImplementation {
  async fn sign_up(&self, req: Req<SignUpRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn sign_in(&self, req: Req<SignInRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
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
