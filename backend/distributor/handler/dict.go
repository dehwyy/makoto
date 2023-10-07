package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
	"github.com/dehwyy/Makoto/backend/distributor/middleware"
	"github.com/dehwyy/Makoto/backend/distributor/pipes"
	dictGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var (
	dict_addr           string
	dict_cancel_timeout = 5 * time.Second
)

func init() {
	dict_host, _ := config.GetOptionByKey("docker_services.dict")
	dict_port, _ := config.GetOptionByKey("server.dict")
	dict_addr = fmt.Sprintf("%s:%s", dict_host, dict_port)
}

// ! Query

func (q *queryResolver) GetWords(ctx context.Context, userId *string) (*model.GetWordsResponse, error) {
	v := middleware.ReadAuthContext(ctx)

	if !v.IsAuth {
		q.log.Errorf("not authenticated %v", *v)
		return nil, nil
	}

	g := rpc()
	g.CreateConnection(dict_addr, q.log)
	defer g.Conn.Close()

	client := dictGrpc.NewDictClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, dict_cancel_timeout)
	defer cancel()

	user_words_id := v.UserId
	if userId != nil {
		user_words_id = *userId
	}

	res, err := client.GetWords(ctx, &dictGrpc.UserId{
		UserId: user_words_id,
	})

	if err != nil {
		return nil, err
	}

	casted_words := pipes.CastWordsGrpcToGraphQL(res.Words)

	return &model.GetWordsResponse{
		Words: casted_words,
		Tokens: &model.Tokens{
			AccessToken:  v.AccessToken,
			RefreshToken: v.RefreshToken,
		},
	}, nil
}

func (q *queryResolver) GetTags(ctx context.Context) (*model.GetTagsResponse, error) {
	v := middleware.ReadAuthContext(ctx)

	if !v.IsAuth {
		q.log.Errorf("not authenticated %v", *v)
		return nil, nil
	}

	g := rpc()
	g.CreateConnection(dict_addr, q.log)
	defer g.Conn.Close()

	client := dictGrpc.NewDictClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, dict_cancel_timeout)
	defer cancel()

	res, err := client.GetTags(ctx, &empty.Empty{})

	if err != nil {
		return nil, err
	}

	return &model.GetTagsResponse{
		Tags: pipes.CastTagsGrpcToGraphQL(res.Tags),
		Tokens: &model.Tokens{
			AccessToken:  v.AccessToken,
			RefreshToken: v.RefreshToken,
		},
	}, nil
}

// ! Mutation

func (m *mutResolver) CreateWord(ctx context.Context, word model.Word) (*model.Tokens, error) {
	v := middleware.ReadAuthContext(ctx)
	if !v.IsAuth {
		return nil, nil
	}

	g := rpc()
	g.CreateConnection(dict_addr, m.log)
	defer g.Conn.Close()

	client := dictGrpc.NewDictClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, dict_cancel_timeout)
	defer cancel()

	payload := &dictGrpc.CreateWord{
		UserId: v.UserId,
		Word: &dictGrpc.Word{
			Word:  word.Word,
			Value: word.Value,
			Extra: word.Extra,
			Tags:  word.Tags,
		},
	}

	res, err := client.CreateNewWord(ctx, payload)
	m.log.Infof("Create word operation: `%v`", res.State)
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		AccessToken:  v.AccessToken,
		RefreshToken: v.RefreshToken,
	}, nil
}

func (m *mutResolver) RemoveWord(ctx context.Context, wordId string) (*model.Tokens, error) {
	v := middleware.ReadAuthContext(ctx)
	if !v.IsAuth {
		return nil, nil
	}

	g := rpc()
	g.CreateConnection(dict_addr, m.log)
	defer g.Conn.Close()

	client := dictGrpc.NewDictClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, dict_cancel_timeout)
	defer cancel()

	// creating payload
	wordIdInt, _ := strconv.Atoi(wordId)

	payload := &dictGrpc.WordId{
		WordId: uint32(wordIdInt),
	}

	res, err := client.RemoveWord(ctx, payload)
	if err != nil {
		return nil, err
	}

	m.log.Infof("Remove word operation: `%v`", res.State)

	return &model.Tokens{
		AccessToken:  v.AccessToken,
		RefreshToken: v.RefreshToken,
	}, nil
}

func (m *mutResolver) EditWord(ctx context.Context, input *model.EditWordInput) (*model.Tokens, error) {
	v := middleware.ReadAuthContext(ctx)
	if !v.IsAuth {
		return nil, nil
	}

	g := rpc()
	g.CreateConnection(dict_addr, m.log)
	defer g.Conn.Close()

	client := dictGrpc.NewDictClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, dict_cancel_timeout)
	defer cancel()

	// creating payload
	wordId, _ := strconv.Atoi(input.WordID)

	payload := &dictGrpc.UpdateWord{
		Id: &dictGrpc.WordId{
			WordId: uint32(wordId),
		},
		Word: &dictGrpc.Word{
			Word:  input.Word.Word,
			Value: input.Word.Value,
			Extra: input.Word.Extra,
			Tags:  input.Word.Tags,
		},
	}

	res, err := client.EditWord(ctx, payload)
	if err != nil {
		return nil, err
	}

	m.log.Infof("Edit word operation: `%v`", res.State)

	return &model.Tokens{
		AccessToken:  v.AccessToken,
		RefreshToken: v.RefreshToken,
	}, nil
}
