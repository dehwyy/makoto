package main

import (
	"net"

	"github.com/dehwyy/Makoto/backend/auth/config"
	"github.com/dehwyy/Makoto/backend/auth/handler"
	"github.com/dehwyy/Makoto/backend/auth/logger"
)

func main() {
	// initialize logger
	l := logger.New()

	// get `port` var from config
	port, isFound := config.GetOptionByKey("server.auth")
	if !isFound {
		l.Fatalf("port variable was not found")
	}

	// make tcp listener listener on port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		l.Fatalf("failed to listen on port %s: %v", port, err)
	}

	// create new grpc server
	srv := handler.NewServer(l)

	l.Infof("Serving on port %s", port)

	// serve it
	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to start server: %v", err)
	}
}
