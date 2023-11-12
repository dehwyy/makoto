package twirp

import (
	"context"
	"fmt"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/user"
	tw "github.com/twitchtv/twirp"
)

type TwirpUserService struct {
	client user.UserRPC
}

func NewUserService(user_service_url string) user.TwirpServer {
	return user.NewUserRPCServer(&TwirpUserService{
		client: services.NewUserService(services.UserService{
			UserServiceUrl: user_service_url,
		}),
	}, tw.WithServerPathPrefix("/user"))
}

func (s *TwirpUserService) CreateUser(ctx context.Context, req *user.CreateUserPayload) (*general.IsSuccess, error) {
	res, err := s.client.CreateUser(ctx, req)
	if err != nil {
		fmt.Printf("error occuried %v", err)
		return nil, err
	}

	fmt.Printf("result %v", res)

	return res, nil
}
