package twirp

import (
	"context"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
)

type TwirpAuthorizationService struct {
	ReadHeader func(context.Context) (string, error)
}

func NewAuthorizationService(args TwirpAuthorizationService) auth.TwirpServer {
	return auth.NewAuthServer(&TwirpAuthorizationService{
		ReadHeader: args.ReadHeader,
	})
}

func (s *TwirpAuthorizationService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	return nil, nil
}

func (s *TwirpAuthorizationService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	return nil, nil
}
