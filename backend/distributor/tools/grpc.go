package tools

import (
	"fmt"

	"github.com/dehwyy/Makoto/backend/distributor/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcConnector struct {
	Conn *grpc.ClientConn
}

func NewGrpcTools() *grpcConnector {
	return &grpcConnector{}
}

// both returns and save to struct Connection, so you can use it as you prefer
func (g *grpcConnector) CreateConnection(addr string, log logger.AppLogger) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if log != nil {
			log.Fatalf("did not connect: %v", err)
		} else {
			fmt.Printf("did not connect: %v", err)
		}
	}

	g.Conn = conn
	return conn
}
