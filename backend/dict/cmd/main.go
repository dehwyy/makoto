package main

import (
	"net"

	"github.com/dehwyy/makoto/backend/dict/config"
	database "github.com/dehwyy/makoto/backend/dict/db"
	"github.com/dehwyy/makoto/backend/dict/handler"
	"github.com/dehwyy/makoto/backend/dict/logger"
)

func main() {
	// init logger
	l := logger.New()

	// Init database
	conn := database.New()
	conn.RunAllMigrations()

	// getting PORT variable
	port, isFound := config.GetOptionByKeyWithFlag("server.dict")
	if !isFound {
		l.Fatalf("port variable was not found")
	}
	// make tcp listener listener on port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		l.Fatalf("failed to listen on port %s: %v", port, err)
	}

	// create new grpc server
	srv := handler.NewServer(l, conn)

	// serve it
	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to start server: %v", err)
	}
}
