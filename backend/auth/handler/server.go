package handler

import (
	"context"

	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/dehwyy/Makoto/backend/auth/tools"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc"
)

type server struct {
	log logger.AppLogger
	auth.UnimplementedUserServer
}

func (s *server) SignUp(ctx context.Context, in *auth.UserCreateRequest) (*auth.UserReply, error) {
	s.log.Infof("Received: %v", in)

	jwt := tools.NewJwt()

	token, _ := jwt.NewRefreshToken(tools.JwtPayload{
		Username: in.Username,
		UserId:   "Flopper",
	})

	err := jwt.ValidateJwtToken(token)
	if err != nil {
		s.log.Debugf("ERROR APPEAR, %v", err)
	}

	s.log.Infof("Token is: %v", token)

	return &auth.UserReply{
		Id: 727,
	}, nil
}

func NewServer(logger logger.AppLogger) *grpc.Server {
	s := grpc.NewServer()
	srv := &server{
		log: logger,
	}

	auth.RegisterUserServer(s, srv)

	return s
}
