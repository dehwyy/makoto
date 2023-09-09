package handler

import (
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/dehwyy/Makoto/backend/auth/service"
	"github.com/dehwyy/Makoto/backend/auth/tools"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc"
)

type server struct {
	log                 logger.AppLogger
	token_service       *service.TokenService
	credentials_service *service.CredentialsService

	auth.UnimplementedUserServer
}

func NewServer(logger logger.AppLogger) *grpc.Server {
	s := grpc.NewServer()
	srv := &server{
		log:                 logger,
		token_service:       service.NewTokenService(tools.NewJwt(), logger),
		credentials_service: service.NewCredentialsService(tools.NewHasher(), logger),
	}

	auth.RegisterUserServer(s, srv)

	return s
}
