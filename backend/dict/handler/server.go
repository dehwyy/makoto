package handler

import (
	dictGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	database "github.com/dehwyy/makoto/backend/dict/db"
	"github.com/dehwyy/makoto/backend/dict/logger"
	"github.com/dehwyy/makoto/backend/dict/service"
	"google.golang.org/grpc"
)

type server struct {
	log           logger.AppLogger
	db            *database.Conn
	tags_service  *service.TagsService
	words_service *service.WordsService

	dictGrpc.UnimplementedDictServer
}

func NewServer(logger logger.AppLogger, db *database.Conn) *grpc.Server {
	s := grpc.NewServer()

	words_service := service.NewWordsService(logger, db)
	tags_service := service.NewTagsService(logger, db)

	srv := &server{
		log:           logger,
		db:            db,
		tags_service:  tags_service,
		words_service: words_service,
	}

	dictGrpc.RegisterDictServer(s, srv)

	return s
}
