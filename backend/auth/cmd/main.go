package main

import (
	"net"

	"github.com/dehwyy/Makoto/backend/auth/config"
	database "github.com/dehwyy/Makoto/backend/auth/db"
	"github.com/dehwyy/Makoto/backend/auth/handler"
	"github.com/dehwyy/Makoto/backend/auth/logger"
)

func main() {
	// initialize logger
	l := logger.New()
	// initiaize database and run migration
	db := database.New(l)
	db.RunAllMigrations()

	// get `port` var from config
	port, isFound := config.GetOptionByKeyWithFlag("server.auth")
	if !isFound {
		l.Fatalf("port variable was not found")
	}

	// make tcp listener listener on port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		l.Fatalf("failed to listen on port %s: %v", port, err)
	}

	// create new grpc server
	srv := handler.NewServer(l, db)

	// serve it
	l.Infof("Serving on port %s", port)

	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to start server: %v", err)
	}
}
