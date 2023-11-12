package services

import (
	"context"
	"strings"

	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/grpc/generated/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	UserServiceUrl string
}

func NewUserService(args UserService) user.UserRPC {
	return &UserService{
		UserServiceUrl: strings.Split(args.UserServiceUrl, "//")[1],
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *user.CreateUserPayload) (*general.IsSuccess, error) {
	user_client, close_fn, err := s.cl(ctx)
	if err != nil {
		return nil, err
	}

	defer close_fn()

	return user_client.CreateUser(ctx, req)
}

func (s *UserService) cl(ctx context.Context) (client user.UserRPCClient, close func(), err error) {
	conn, err := grpc.Dial(s.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	close = func() {
		conn.Close()
	}

	return user.NewUserRPCClient(conn), close, nil
}
