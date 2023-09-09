package handler

import (
	"github.com/dehwyy/Makoto/backend/auth/logger"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc"
)

type server struct {
	log logger.AppLogger
	auth.UnimplementedUserServer
}

func NewServer(logger logger.AppLogger) *grpc.Server {
	s := grpc.NewServer()
	srv := &server{
		log: logger,
	}

	auth.RegisterUserServer(s, srv)

	return s
}
