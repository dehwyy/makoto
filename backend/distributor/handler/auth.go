package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
	"github.com/dehwyy/Makoto/backend/distributor/middleware"
	"github.com/dehwyy/Makoto/backend/distributor/tools"
	authGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
)

var (
	authAddr string
)

func init() {
	authHost, _ := config.GetOptionByKey("docker_services.auth")
	authPort, _ := config.GetOptionByKey("server.auth")
	authAddr = fmt.Sprintf("%s:%s", authHost, authPort)
}

func (m *mutResolver) SignUp(ctx context.Context, input *model.SignUpInput) (*model.UserAuthResponse, error) {
	// TODO: shouldn't be used in /auth microservice
	v, _ := middleware.ReadAuthContext(ctx)
	m.log.Debugf("Context value is %v", *v)

	// init and call to create new grpc connection
	g := tools.NewGrpcTools()
	// read desc of func
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	payload := &authGrpc.UserSignUpRequest{
		Username: input.Username,
		Password: input.Password,
		Question: input.Question,
		Answer:   input.Answer,
	}

	res, err := cl.SignUp(ctx, payload)
	if err != nil {
		m.log.Errorf("Error calling SignUp: %v", err)
	}

	return &model.UserAuthResponse{
		UserID: res.UserId,
		Tokens: &model.Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

func (q *queryResolver) GetQuestion(ctx context.Context, username string) (string, error) {
	return "", nil
}
