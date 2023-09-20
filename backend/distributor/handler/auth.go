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
	authAddr            string
	rpc                 = tools.NewGrpcTools
	auth_cancel_timeout = 5 * time.Second
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
	g := rpc()
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
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
	g := rpc()
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

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

func (m *mutResolver) SignOut(ctx context.Context) (*model.Status, error) {
	g := rpc()
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, auth_cancel_timeout)
	defer cancel()

	userId := middleware.ReadAuthContext(ctx).UserId

	// this function literally returns Nil xd ( so just ignore 1st value as it would be empty empty), but error anyway could appear
	_, err := cl.SignOut(ctx, &authGrpc.UserSignOutRequest{
		UserId: userId,
	})

	if err != nil {
		// by default the value "isOk" should be false. Isn't it?)
		return new(model.Status), err
	}

	return &model.Status{
		IsOk: true,
	}, nil
}

func (m *mutResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.UserAuthResponse, error) {
	g := rpc()
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, auth_cancel_timeout)
	defer cancel()

	userId := middleware.ReadAuthContext(ctx).UserId

	res, err := cl.ChangePassword(ctx, &authGrpc.UserChangePasswordRequest{
		UserId:      userId,
		OldPassword: input.OldPassword,
		NewPassword: input.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		UserID: userId,
		Tokens: &model.Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

func (m *mutResolver) ChangePasswordByAnswer(ctx context.Context, input model.ChangePasswordByAnswerInput) (*model.UserAuthResponse, error) {
	g := rpc()
	g.CreateConnection(authAddr, m.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, auth_cancel_timeout)
	defer cancel()

	userId := middleware.ReadAuthContext(ctx).UserId

	res, err := cl.ChangePasswordByAnswer(ctx, &authGrpc.UserSendAnswerAndChangePasswordRequest{
		UserId:      userId,
		Answer:      input.Answer,
		NewPassword: input.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		UserID: userId,
		Tokens: &model.Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

// ! Queries

func (q *queryResolver) GetUserByUsername(ctx context.Context, input model.GetUserByUsernameInput) (*model.UserResponse, error) {
	return nil, nil
}

func (q *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDInput) (*model.UserResponse, error) {
	auth_ctx := middleware.ReadAuthContext(ctx)
	if !auth_ctx.IsAuth {
		return nil, fmt.Errorf("404:user is not authenticated")
	}

	g := rpc()
	g.CreateConnection(authAddr, nil)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := cl.GetUserById(ctx, &authGrpc.UserGetRequest{
		UserId: auth_ctx.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		Username: res.Username,
		Auth:     authf.createAuthResponseByCtx(ctx),
	}, nil
}

func (q *queryResolver) GetQuestion(ctx context.Context) (*model.UserQuestionResponse, error) {
	g := rpc()
	g.CreateConnection(authAddr, q.log)
	defer g.Conn.Close()

	cl := authGrpc.NewUserClient(g.Conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	auth_ctx := middleware.ReadAuthContext(ctx)
	if !auth_ctx.IsAuth {
		return nil, fmt.Errorf("user is not authenticated")
	}

	payload := &authGrpc.UserGetQuestionRequest{
		UserId: auth_ctx.UserId,
	}

	res, err := cl.GetQuestion(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &model.UserQuestionResponse{
		Question: res.Question,
		Auth:     authf.createAuthResponseByCtx(ctx),
	}, nil
}
