package twirp

import (
	"context"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
)

type Server struct{}

func NewTwirpServer() auth.TwirpServer {
	return auth.NewAuthServer(&Server{})
}

func (s *Server) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	return nil, nil
}
