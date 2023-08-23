package handler

import (
	"github.com/dehwyy/Makoto/backend/distributor/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func grpcConnection(addr string, log logger.AppLogger) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}
