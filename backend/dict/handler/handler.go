package handler

import (
	"context"

	rpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	"github.com/dehwyy/makoto/backend/dict/pipes"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *server) GetWords(ctx context.Context, in *rpc.UserId) (*rpc.Words, error) {
	words, err := s.words_service.GetWords(in.UserId)
	if err != nil {
		return nil, err
	}

	casted_words := pipes.CastWordsToGrpcOutput(words)

	return &rpc.Words{
		Words: casted_words,
	}, nil
}

func (s *server) GetTags(context.Context, *empty.Empty) (*rpc.TagsResponse, error) {
	tags := s.tags_service.GetAllTags()

	return &rpc.TagsResponse{
		Tags: pipes.CastTagsToGrpcOutput(tags),
	}, nil
}

func (s *server) CreateNewWord(ctx context.Context, in *rpc.CreateWord) (*rpc.Status, error) {

	// getting tags from Array<string>
	tags := s.tags_service.TagsFromStringArray(in.Word.Tags)

	err := s.words_service.CreateWord(in.UserId, in.Word.Word, in.Word.Value, in.Word.Extra, tags)

	return &rpc.Status{
		State: err == nil,
	}, err
}

func (s *server) RemoveWord(ctx context.Context, in *rpc.WordId) (*rpc.Status, error) {

	err := s.words_service.RemoveWord(in.WordId)

	return &rpc.Status{
		State: err == nil,
	}, err
}

func (s *server) EditWord(ctx context.Context, in *rpc.UpdateWord) (*rpc.Status, error) {

	// getting tags from Array<string>
	tags := s.tags_service.TagsFromStringArray(in.Word.Tags)

	err := s.words_service.EditWord(in.Id.WordId, in.Word.Word, in.Word.Value, in.Word.Extra, tags)

	return &rpc.Status{
		State: err == nil,
	}, err
}
