package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/apps/gateway/twirp/internal/middleware"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	"github.com/golang/protobuf/ptypes/empty"
	tw "github.com/twitchtv/twirp"
)

type Empty = empty.Empty

type TwirpHashmapService struct {
	ReadAuthorizationData func(context.Context) middleware.AuthCredentialsGranted

	client hashmap.HashmapRPC
}

func NewHashmapService(hashmap_service_url string, args TwirpHashmapService) hashmap.TwirpServer {
	return hashmap.NewHashmapRPCServer(&TwirpHashmapService{
		ReadAuthorizationData: args.ReadAuthorizationData,

		client: services.NewHashmapService(services.HashmapService{
			HashmapServiceUrl: hashmap_service_url,
		}),
	}, tw.WithServerPathPrefix("/hashmap"))
}

func (s *TwirpHashmapService) GetItems(ctx context.Context, req *hashmap.GetItemsPayload) (*hashmap.GetItemsResponse, error) {
	userId := s.ReadAuthorizationData(ctx).UserId()
	req.UserId = userId

	response, err := s.client.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) GetTags(ctx context.Context, req *general.UserId) (*hashmap.GetTagsResponse, error) {

	userId := s.ReadAuthorizationData(ctx).UserId()

	new_req := &general.UserId{
		UserId: userId,
	}

	response, err := s.client.GetTags(ctx, new_req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) CreateItem(ctx context.Context, req *hashmap.CreateItemPayload) (*hashmap.CreateItemResponse, error) {
	userId := s.ReadAuthorizationData(ctx).UserId()

	new_req := &hashmap.CreateItemPayload{
		UserId: userId,
		Key:    req.Key,
		Value:  req.Value,
		Extra:  req.Extra,
		Tags:   req.Tags,
	}

	response, err := s.client.CreateItem(ctx, new_req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) RemoveItem(ctx context.Context, req *hashmap.RemoveItemPayload) (*general.IsSuccess, error) {

	userId := s.ReadAuthorizationData(ctx).UserId()

	new_req := &hashmap.RemoveItemPayload{
		UserId: userId,
		ItemId: req.ItemId,
	}

	response, err := s.client.RemoveItem(ctx, new_req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpHashmapService) EditItem(ctx context.Context, req *hashmap.EditItemPayload) (*general.IsSuccess, error) {

	userId := s.ReadAuthorizationData(ctx).UserId()

	new_req := &hashmap.EditItemPayload{
		UserId: userId,
		ItemId: req.ItemId,
		Key:    req.Key,
		Value:  req.Value,
		Extra:  req.Extra,
		Tags:   req.Tags,
	}

	response, err := s.client.EditItem(ctx, new_req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
