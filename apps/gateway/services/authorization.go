package services

import (
	"context"
	"net/http"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
)

type AuthorizationService struct {
	AuthorizationServiceUrl string
}

func NewAuthorizationService(args AuthorizationService) auth.AuthRPC {
	return &AuthorizationService{
		AuthorizationServiceUrl: args.AuthorizationServiceUrl,
	}
}

func (s *AuthorizationService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	auth_client := s.cl(ctx)
	return auth_client.SignUp(ctx, req)
}

func (s *AuthorizationService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	auth_client := s.cl(ctx)
	return auth_client.SignIn(ctx, req)
}

func (s *AuthorizationService) IsUniqueUsername(ctx context.Context, req *auth.IsUniqueUsernamePayload) (*auth.IsUnique, error) {
	auth_client := s.cl(ctx)
	return auth_client.IsUniqueUsername(ctx, req)
}

func (s *AuthorizationService) VerifyUserEmail(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	auth_client := s.cl(ctx)
	return auth_client.VerifyUserEmail(ctx, req)
}

func (s *AuthorizationService) ChangePassword(ctx context.Context, req *auth.ChangePasswordPayload) (*general.IsSuccess, error) {
	auth_client := s.cl(ctx)
	return auth_client.ChangePassword(ctx, req)
}

func (s *AuthorizationService) Logout(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	auth_client := s.cl(ctx)
	return auth_client.Logout(ctx, req)
}

func (s *AuthorizationService) cl(ctx context.Context) auth.AuthRPC {
	return auth.NewAuthRPCProtobufClient(s.AuthorizationServiceUrl, http.DefaultClient)
}
