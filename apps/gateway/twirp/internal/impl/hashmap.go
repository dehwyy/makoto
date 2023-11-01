package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/golang/protobuf/ptypes/empty"
)

type Empty = empty.Empty

type TwirpHashmapService struct {
	ReadAuthorizationData func(context.Context) (userId, token string)

	client hashmap.HashmapRPC
}

func NewHashmapService(hashmap_service_url string, args TwirpHashmapService) hashmap.TwirpServer {
	return hashmap.NewHashmapRPCServer(&TwirpHashmapService{
		ReadAuthorizationData: args.ReadAuthorizationData,

		client: services.NewHashmapService(services.HashmapService{
			HashmapServiceUrl: hashmap_service_url,
		}),
	})
}

func (s *TwirpHashmapService) GetItems(ctx context.Context, req *hashmap.GetItemsPayload) (*hashmap.GetItemsResponse, error) {
	response, err := s.client.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) GetTags(ctx context.Context, req *general.UserId) (*hashmap.GetTagsResponse, error) {
	response, err := s.client.GetTags(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) CreateItem(ctx context.Context, req *hashmap.CreateItemPayload) (*hashmap.CreateItemResponse, error) {
	response, err := s.client.CreateItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) RemoveItem(ctx context.Context, req *hashmap.RemoveItemPayload) (*general.IsSuccess, error) {
	response, err := s.client.RemoveItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) EditItem(ctx context.Context, req *hashmap.EditItemPayload) (*general.IsSuccess, error) {
	response, err := s.client.EditItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
