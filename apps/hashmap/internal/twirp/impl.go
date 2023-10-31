package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/hashmap/internal/pipes"
	"github.com/dehwyy/makoto/apps/hashmap/internal/repository"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	tw "github.com/twitchtv/twirp"
	"gorm.io/gorm"
)

type Server struct {
	// repos
	tags_repository  *repository.TagsRepository
	items_repository *repository.ItemsRepository

	//
	l logger.Logger
}

type Empty = empty.Empty

var (
	InvalidUserIdError = tw.PermissionDenied.Error("invalid user id")
)

func NewTwirpServer(db *gorm.DB, l logger.Logger) hashmap.TwirpServer {
	return hashmap.NewHashmapRPCServer(
		&Server{
			//
			tags_repository:  repository.NewTagsRepository(db, l),
			items_repository: repository.NewItemsRepository(db, l),
			//
			l: l,
		},
	)
}

func (s *Server) GetItems(ctx context.Context, req *hashmap.GetItemsPayload) (*hashmap.GetItemsResponse, error) {
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, InvalidUserIdError
	}

	items, err := s.items_repository.GetItems(user_id)
	if err != nil {
		return nil, tw.InternalErrorf("failed to get items: %v", err.Error())
	}

	return &hashmap.GetItemsResponse{
		Items: pipes.ToRpcItems(items),
	}, nil
}

func (s *Server) GetTags(ctx context.Context, req *general.UserId) (*hashmap.GetTagsResponse, error) {
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, InvalidUserIdError
	}

	tags, err := s.tags_repository.GetAllTags(user_id)
	if err != nil {
		return nil, tw.InternalErrorf("failed to get tags: %v", err.Error())
	}

	return &hashmap.GetTagsResponse{
		Tags: pipes.ToRpcTags(tags),
	}, nil
}

func (s *Server) CreateItem(ctx context.Context, req *hashmap.CreateItemPayload) (*general.IsSuccess, error) {
	tags := s.tags_repository.TagsFromStringArray(req.Tags)

	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, tw.InvalidArgumentError("invalid user id (not uuid)", err.Error())
	}

	item_id, err := s.items_repository.CreateItem(user_id, req.Key, req.Value, req.Extra, tags)
	if err != nil {
		return nil, tw.InternalErrorf("failed to create item: %v", err.Error())
	}

	s.l.Debugf("item created: %v", item_id)

	return &general.IsSuccess{
		IsSuccess: true,
	}, nil
}

func (s *Server) RemoveItem(ctx context.Context, req *hashmap.RemoveItemPayload) (*general.IsSuccess, error) {
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, InvalidUserIdError
	}

	err = s.items_repository.RemoveItem(user_id, uint32(req.ItemId))
	if err != nil {
		return nil, tw.InternalErrorf("failed to remove item: %v", err.Error())
	}

	return &general.IsSuccess{
		IsSuccess: true,
	}, nil
}

func (s *Server) EditItem(ctx context.Context, req *hashmap.EditItemPayload) (*general.IsSuccess, error) {
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, InvalidUserIdError
	}

	// generating tags (in db) or reading (from db)
	tags := s.tags_repository.TagsFromStringArray(req.Tags)

	err = s.items_repository.EditItem(user_id, req.ItemId, req.Key, req.Value, req.Extra, tags)
	if err != nil {
		return nil, tw.InternalErrorf("failed to edit item: %v", err.Error())
	}

	return &general.IsSuccess{
		IsSuccess: true,
	}, nil
}
