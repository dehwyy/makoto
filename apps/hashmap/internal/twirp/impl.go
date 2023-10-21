package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/hashmap/internal/pipes"
	"github.com/dehwyy/makoto/apps/hashmap/internal/repository"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/middleware"
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
	return hashmap.NewHashmapServer(
		&Server{
			//
			tags_repository:  repository.NewTagsRepository(db, l),
			items_repository: repository.NewItemsRepository(db, l),
			//
			l: l,
		},
	)
}

func (s *Server) GetItems(ctx context.Context, req *hashmap.UserId) (*hashmap.Items, error) {
	user_id, err := s.getUUIDFromContext(ctx)
	if err != nil {
		return nil, InvalidUserIdError
	}

	items, err := s.items_repository.GetItems(user_id)
	if err != nil {
		return nil, tw.InternalErrorf("failed to get items: %v", err.Error())
	}

	return &hashmap.Items{
		Items: pipes.ItemsDb2Grpc(items),
	}, nil
}

// No authorization required
func (s *Server) GetTags(ctx context.Context, req *hashmap.UserId) (*hashmap.TagsResponse, error) {
	user_id, err := s.getUUIDFromContext(ctx)
	if err != nil {
		return nil, InvalidUserIdError
	}

	if req.UserId != "" {
		user_id, err = uuid.Parse(req.UserId)
		if err != nil {
			return nil, tw.InvalidArgumentError("invalid user id", err.Error())
		}
	}

	tags, err := s.tags_repository.GetAllTags(user_id)
	if err != nil {
		return nil, tw.InternalErrorf("failed to get tags: %v", err.Error())
	}

	return &hashmap.TagsResponse{
		Tags: pipes.TagsFromDb2Grpc(tags),
	}, nil
}

func (s *Server) CreateItem(ctx context.Context, req *hashmap.Item) (*hashmap.ItemId, error) {
	tags := s.tags_repository.TagsFromStringArray(req.Tags)
	user_id, err := s.getUUIDFromContext(ctx)
	if err != nil {
		return nil, tw.InvalidArgumentError("invalid user id (not uuid)", err.Error())
	}

	item_id, err := s.items_repository.CreateItem(user_id, req.Key, req.Value, req.Extra, tags)
	if err != nil {
		return nil, tw.InternalErrorf("failed to create item: %v", err.Error())
	}

	s.l.Debugf("item created: %v", item_id)

	return &hashmap.ItemId{
		ItemId: item_id,
	}, nil
}

func (s *Server) RemoveItem(ctx context.Context, req *hashmap.ItemId) (*Empty, error) {
	user_id, err := s.getUUIDFromContext(ctx)
	if err != nil {
		return nil, InvalidUserIdError
	}

	err = s.items_repository.RemoveItem(user_id, uint32(req.ItemId))
	if err != nil {
		return nil, tw.InternalErrorf("failed to remove item: %v", err.Error())
	}

	return &Empty{}, nil
}

func (s *Server) EditItem(ctx context.Context, req *hashmap.UpdateItem) (*Empty, error) {
	user_id, err := s.getUUIDFromContext(ctx)
	if err != nil {
		return nil, InvalidUserIdError
	}

	// generating tags (in db) or reading (from db)
	tags := s.tags_repository.TagsFromStringArray(req.Item.Tags)

	err = s.items_repository.EditItem(user_id, req.Id.ItemId, req.Item.Key, req.Item.Value, req.Item.Extra, tags)
	if err != nil {
		return nil, tw.InternalErrorf("failed to edit item: %v", err.Error())
	}

	return &Empty{}, nil
}

func (s *Server) getUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	return uuid.Parse(middleware.AuthorizationMiddlewareRead(ctx))
}
