package handler

import (
	dictGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	database "github.com/dehwyy/makoto/backend/dict/db"
	"github.com/dehwyy/makoto/backend/dict/logger"
	"google.golang.org/grpc"
)

type server struct {
	log logger.AppLogger
	db  *database.Conn

	dictGrpc.UnimplementedDictServer
}

func NewServer(logger logger.AppLogger, db *database.Conn) *grpc.Server {
	s := grpc.NewServer()
	srv := &server{
		log: logger,
		db:  db,
	}

	dictGrpc.RegisterDictServer(s, srv)

	return s
}
