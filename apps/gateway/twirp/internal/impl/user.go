package twirp

import (
	"context"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/apps/gateway/twirp/internal/middleware"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/user"
	tw "github.com/twitchtv/twirp"
)

type TwirpUserService struct {
	ReadAuthorizationData func(context.Context) middleware.AuthCredentialsGranted
	client                user.UserRPC
}

func NewUserService(user_service_url string, args TwirpUserService) user.TwirpServer {
	return user.NewUserRPCServer(&TwirpUserService{
		ReadAuthorizationData: args.ReadAuthorizationData,
		client: services.NewUserService(services.UserService{
			UserServiceUrl: user_service_url,
		}),
	}, tw.WithServerPathPrefix("/user"))
}

// ! SHOULN'T BE CALLED DIRECTLY
func (s *TwirpUserService) CreateUser(ctx context.Context, req *user.CreateUserPayload) (*general.IsSuccess, error) {
	return nil, nil
}

func (s *TwirpUserService) GetUser(ctx context.Context, req *user.GetUserPayload) (*user.GetUserResponse, error) {
	userId := s.ReadAuthorizationData(ctx).UserId()

	if req.UserId == "" {
		req.UserId = userId
	}

	res, err := s.client.GetUser(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *TwirpUserService) UpdateUser(ctx context.Context, req *user.UpdateUserPayload) (*general.IsSuccess, error) {
	userId := s.ReadAuthorizationData(ctx).UserId()

	if req.UserId == "" {
		req.UserId = userId
	}

	res, err := s.client.UpdateUser(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
