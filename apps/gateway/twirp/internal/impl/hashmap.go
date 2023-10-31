package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/golang/protobuf/ptypes/empty"
)

type Empty = empty.Empty

type TwirpHashmapService struct {
	ReadAuthorizationData func(context.Context) (userId, token string)

	client hashmap.Hashmap
}

func NewHashmapService(hashmap_service_url string, args TwirpHashmapService) hashmap.TwirpServer {
	return hashmap.NewHashmapServer(&TwirpHashmapService{
		ReadAuthorizationData: args.ReadAuthorizationData,

		client: services.NewHashmapService(services.HashmapService{
			HashmapServiceUrl: hashmap_service_url,
		}),
	})
}

func (s *TwirpHashmapService) GetItems(ctx context.Context, req *hashmap.UserId) (*hashmap.Items, error) {
	response, err := s.client.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) GetTags(ctx context.Context, req *hashmap.UserId) (*hashmap.TagsResponse, error) {
	response, err := s.client.GetTags(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) CreateItem(ctx context.Context, req *hashmap.Item) (*hashmap.ItemId, error) {
	response, err := s.client.CreateItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) RemoveItem(ctx context.Context, req *hashmap.ItemId) (*Empty, error) {
	response, err := s.client.RemoveItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) EditItem(ctx context.Context, req *hashmap.UpdateItem) (*Empty, error) {
	response, err := s.client.EditItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
