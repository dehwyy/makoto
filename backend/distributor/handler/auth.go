package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
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

func (m *mutResolver) SignUp(ctx context.Context, input *model.SignUpInput) (string, error) {
	fmt.Println(authAddr)

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
		m.log.Errorf("Error calling SignUp: %v", err)
	}

	m.log.Debugf("Received: %v", res.Id)

	return strconv.Itoa(int(res.Id)), nil
}

func (q *queryResolver) GetQuestion(ctx context.Context, username string) (string, error) {
	return "", nil
}
