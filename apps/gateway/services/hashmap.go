package services

import (
	"context"
	"net/http"

	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/golang/protobuf/ptypes/empty"
)

type HashmapService struct {
	HashmapServiceUrl string
}

func NewHashmapService(args HashmapService) hashmap.Hashmap {
	return &HashmapService{
		HashmapServiceUrl: args.HashmapServiceUrl,
	}
}

func (s *HashmapService) GetItems(ctx context.Context, req *hashmap.UserId) (*hashmap.Items, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.GetItems(ctx, req)
}

func (s *HashmapService) GetTags(ctx context.Context, req *hashmap.UserId) (*hashmap.TagsResponse, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.GetTags(ctx, req)
}

func (s *HashmapService) CreateItem(ctx context.Context, req *hashmap.Item) (*hashmap.ItemId, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.CreateItem(ctx, req)
}

func (s *HashmapService) RemoveItem(ctx context.Context, req *hashmap.ItemId) (*empty.Empty, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.RemoveItem(ctx, req)
}

func (s *HashmapService) EditItem(ctx context.Context, req *hashmap.UpdateItem) (*empty.Empty, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.EditItem(ctx, req)
}

func (s *HashmapService) cl(ctx context.Context) hashmap.Hashmap {
	return hashmap.NewHashmapProtobufClient(s.HashmapServiceUrl, http.DefaultClient)
}
