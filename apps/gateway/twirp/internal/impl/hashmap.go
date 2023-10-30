package twirp

import (
	"context"

	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/golang/protobuf/ptypes/empty"
)

type Empty = empty.Empty

type TwirpHashmapService struct {
	ReadAuthorizationData func(context.Context) (userId, token string, err error)
}

func NewHashmapService(args TwirpHashmapService) hashmap.TwirpServer {
	return hashmap.NewHashmapServer(&TwirpHashmapService{
		ReadAuthorizationData: args.ReadAuthorizationData,
	})
}

func (s *TwirpHashmapService) GetItems(context.Context, *hashmap.UserId) (*hashmap.Items, error) {
	return nil, nil
}

func (s *TwirpHashmapService) GetTags(context.Context, *hashmap.UserId) (*hashmap.TagsResponse, error) {
	return nil, nil
}

func (s *TwirpHashmapService) CreateItem(context.Context, *hashmap.Item) (*hashmap.ItemId, error) {
	return nil, nil
}

func (s *TwirpHashmapService) RemoveItem(context.Context, *hashmap.ItemId) (*Empty, error) {
	return nil, nil
}

func (s *TwirpHashmapService) EditItem(context.Context, *hashmap.UpdateItem) (*Empty, error) {
	return nil, nil
}
