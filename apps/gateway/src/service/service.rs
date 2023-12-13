use makoto_grpc::pkg::api_auth::{SignUpRequest, SignInRequest, AuthorizationResponse,
  SignInOauthRequest, SignInTokenRequest,
  SignOutRequest, ConfirmMailByCodeRequest, ProceedToUpdatePasswordResponse,
  ProceedToRecoverPasswordRequest, ProceedToUpdatePasswordRequest, SubmitNewPasswordByCodeRequest,
  IsEmailAvailableRequest, IsUsernameAvailableRequest};

use makoto_grpc::pkg::api::api_rpc_server::ApiRpc;
use makoto_grpc::pkg::general::BoolStatus;

use makoto_grpc::Result as GrpcResult;
use tonic::Request as Req;

#[derive(Default)]
pub struct ApiRpcServiceImplementation;

impl ApiRpcServiceImplementation {
  pub fn new() -> Self {
    Self {
    }
  }
}

#[tonic::async_trait]
impl ApiRpc for ApiRpcServiceImplementation {
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

  async fn confirm_mail_by_code(&self, req: Req<ConfirmMailByCodeRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn proceed_to_update_password(&self, req: Req<ProceedToUpdatePasswordRequest>) -> GrpcResult<ProceedToUpdatePasswordResponse> {
    todo!()
  }

  async fn proceed_to_recover_password(&self, req: Req<ProceedToRecoverPasswordRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }

  async fn submit_new_password_by_code(&self, req: Req<SubmitNewPasswordByCodeRequest>) -> GrpcResult<AuthorizationResponse> {
    todo!()
  }

  async fn is_email_available(&self, req: Req<IsEmailAvailableRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }

  async fn is_username_available(&self, req: Req<IsUsernameAvailableRequest>) -> GrpcResult<BoolStatus> {
    todo!()
  }
}
