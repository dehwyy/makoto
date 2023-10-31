package twirp

import (
	"context"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
)

type TwirpAuthorizationService struct {
	ReadHeader func(context.Context) (string, error)
}

func NewAuthorizationService(args TwirpAuthorizationService) auth.TwirpServer {
	return auth.NewAuthRPCServer(&TwirpAuthorizationService{
		ReadHeader: args.ReadHeader,
	})
}

func (s *TwirpAuthorizationService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) IsUniqueUsername(ctx context.Context, req *auth.IsUniqueUsernamePayload) (*auth.IsUnique, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) VerifyUserEmail(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) ChangePassword(ctx context.Context, req *auth.ChangePasswordPayload) (*general.IsSuccess, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) Logout(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	return nil, nil
}
