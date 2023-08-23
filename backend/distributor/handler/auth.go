package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	authGrpc "github.com/dehwyy/Makoto/backend/auth/grpc"
	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
)

var (
	authAddr string
)

func init() {
	authPort, _ := config.GetOptionByKey("server.auth")
	authAddr = fmt.Sprintf("localhost:%s", authPort)
}

func (m *mutResolver) SignUp(ctx context.Context, input *model.SignUpInput) (string, error) {
	conn := grpcConnection(authAddr, m.log)
	defer conn.Close()

	cl := authGrpc.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	payload := &authGrpc.UserCreateRequest{
		Username: input.Username,
		Password: input.Password,
		Question: input.Question,
		Answer:   input.Answer,
	}

	res, err := cl.SignUp(ctx, payload)
	if err != nil {
		m.log.Fatalf("Error calling SignUp: %v", err)
	}

	m.log.Debugf("Received: %v", res.Id)

	return strconv.Itoa(int(res.Id)), nil
}

func (q *queryResolver) GetQuestion(ctx context.Context, username string) (string, error) {
	return "", nil
}
