package handler

import (
	"context"

	rpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
)

func (s *server) GetWords(ctx context.Context, in *rpc.UserId) (*rpc.Words, error) {
	return nil, nil
}

func (s *server) CreateNewWord(ctx context.Context, in *rpc.CreateWord) (*rpc.Status, error) {
	return nil, nil
}

func (s *server) RemoveWord(ctx context.Context, in *rpc.WordId) (*rpc.Status, error) {
	return nil, nil
}

func (s *server) EditWord(ctx context.Context, in *rpc.UpdateWord) (*rpc.Status, error) {
	return nil, nil
}
