package services

import (
	"context"
	"net/http"

	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
)

type HashmapService struct {
	HashmapServiceUrl string
}

func NewHashmapService(args HashmapService) hashmap.HashmapRPC {
	return &HashmapService{
		HashmapServiceUrl: args.HashmapServiceUrl,
	}
}

func (s *HashmapService) GetItems(ctx context.Context, req *hashmap.GetItemsPayload) (*hashmap.GetItemsResponse, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.GetItems(ctx, req)
}

func (s *HashmapService) GetTags(ctx context.Context, req *general.UserId) (*hashmap.GetTagsResponse, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.GetTags(ctx, req)
}

func (s *HashmapService) CreateItem(ctx context.Context, req *hashmap.CreateItemPayload) (*general.IsSuccess, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.CreateItem(ctx, req)
}

func (s *HashmapService) RemoveItem(ctx context.Context, req *hashmap.RemoveItemPayload) (*general.IsSuccess, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.RemoveItem(ctx, req)
}

func (s *HashmapService) EditItem(ctx context.Context, req *hashmap.EditItemPayload) (*general.IsSuccess, error) {
	hashmap_client := s.cl(ctx)
	return hashmap_client.EditItem(ctx, req)
}

func (s *HashmapService) cl(ctx context.Context) hashmap.HashmapRPC {
	return hashmap.NewHashmapRPCProtobufClient(s.HashmapServiceUrl, http.DefaultClient)
}
