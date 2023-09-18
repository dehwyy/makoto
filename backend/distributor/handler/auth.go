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
	rpc      = tools.NewGrpcTools()
)

func init() {
	authHost, _ := config.GetOptionByKey("docker_services.auth")
	authPort, _ := config.GetOptionByKey("server.auth")
	authAddr = fmt.Sprintf("%s:%s", authHost, authPort)
}

//! Mutations

// no token
func (m *mutResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.UserAuthResponse, error) {
	// TODO: shouldn't be used in /auth microservice
	v := middleware.ReadAuthContext(ctx)
	m.log.Debugf("Context value is %v", *v)

	// read desc of func
	rpc.CreateConnection(authAddr, m.log)
	defer rpc.Conn.Close()

	cl := authGrpc.NewUserClient(rpc.Conn)

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
		return nil, err
	}

	return &model.UserAuthResponse{
		UserID: res.UserId,
		Tokens: &model.Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

// no token
func (m *mutResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.UserAuthResponse, error) {
	rpc.CreateConnection(authAddr, m.log)
	defer rpc.Conn.Close()

	cl := authGrpc.NewUserClient(rpc.Conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	payload := &authGrpc.UserSignInRequest{
		Username: input.Username,
		Password: input.Password,
	}

	res, err := cl.SignIn(ctx, payload)
	if err != nil {
		m.log.Errorf("Error calling SignIn: %v", err)
		return nil, err
	}

	return &model.UserAuthResponse{
		UserID: res.UserId,
		Tokens: &model.Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

func (m *mutResolver) SignOut(ctx context.Context, input model.Token) (bool, error) {
	return false, nil
}

func (m *mutResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.UserAuthResponse, error) {
	return nil, nil
}

func (m *mutResolver) ChangePasswordByAnswer(ctx context.Context, input model.ChangePasswordByAnswerInput) (*model.UserAuthResponse, error) {
	return nil, nil
}

// ! Queries

func (q *queryResolver) GetUserByUsername(ctx context.Context, input model.GetUserByUsernameInput) (*model.UserResponse, error) {
	return nil, nil
}

func (q *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDInput) (*model.UserResponse, error) {
	return nil, nil
}

func (q *queryResolver) GetQuestion(ctx context.Context, input model.Token) (*model.UserQuestionResponse, error) {
	return nil, nil
}
