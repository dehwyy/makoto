package handler

import (
	"context"

	"github.com/dehwyy/Makoto/backend/auth/service"
	"github.com/dehwyy/Makoto/backend/auth/tools"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
)

func (s *server) SignUp(ctx context.Context, in *auth.UserSignUpRequest) (*auth.UserResponse, error) {
	service.NewTokenService(tools.NewJwt())
}
