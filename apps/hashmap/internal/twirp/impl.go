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

// TODO: twirp error

func (s *Server) GetItems(ctx context.Context, req *hashmap.UserId) (*hashmap.Items, error) {
	user_id := middleware.AuthorizationMiddlewareRead(ctx)
	if req.UserId != "" {
		user_id = req.UserId
	}

	parsed_user_id, err := uuid.Parse(user_id)
	if err != nil {
		return nil, tw.InvalidArgumentError("invalid user id (not uuid)", err.Error())
	}

	items, err := s.items_repository.GetItems(parsed_user_id)

	return &hashmap.Items{
		Items: pipes.ItemsDb2Grpc(items),
	}, err
}

// ? NO AUTH REQUIRED
// ? OR SHOULD BE?

func (s *Server) GetTags(ctx context.Context, req *Empty) (*hashmap.TagsResponse, error) {
	tags, err := s.tags_repository.GetAllTags()

	return &hashmap.TagsResponse{
		Tags: pipes.TagsFromDb2Grpc(tags),
	}, err
}

func (s *Server) CreateItem(ctx context.Context, req *hashmap.Item) (*Empty, error) {
	tags := s.tags_repository.TagsFromStringArray(req.Tags)
	parsed_user_id, err := uuid.Parse(middleware.AuthorizationMiddlewareRead(ctx))
	if err != nil {
		return nil, tw.InvalidArgumentError("invalid user id (not uuid)", err.Error())
	}

	return &Empty{}, s.items_repository.CreateItem(parsed_user_id, req.Key, req.Value, req.Extra, tags)
}

// TODO: clarify that Item belongs to User (by authorization-header (userId))
func (s *Server) RemoveItem(ctx context.Context, req *hashmap.ItemId) (*Empty, error) {

	return &Empty{}, s.items_repository.RemoveItem(uint32(req.ItemId))
}

// TODO: same as above
func (s *Server) EditItem(ctx context.Context, req *hashmap.UpdateItem) (*Empty, error) {
	tags := s.tags_repository.TagsFromStringArray(req.Item.Tags)

	return &Empty{}, s.items_repository.EditItem(req.Id.ItemId, req.Item.Key, req.Item.Value, req.Item.Extra, tags)
}
