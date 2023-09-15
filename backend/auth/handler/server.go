package handler

import (
	database "github.com/dehwyy/Makoto/backend/auth/db"
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/dehwyy/Makoto/backend/auth/service"
	"github.com/dehwyy/Makoto/backend/auth/tools"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc"
)

type server struct {
	log                 logger.AppLogger
	db                  *database.Conn
	token_service       *service.TokenService
	credentials_service *service.CredentialsService

	auth.UnimplementedUserServer
}

func NewServer(logger logger.AppLogger, db *database.Conn) *grpc.Server {
	s := grpc.NewServer()
	srv := &server{
		log:                 logger,
		db:                  db,
		token_service:       service.NewTokenService(tools.NewJwt(), db.DB, logger),
		credentials_service: service.NewCredentialsService(tools.NewHasher(), db.DB, logger),
	}

	auth.RegisterUserServer(s, srv)

	return s
}
