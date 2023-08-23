package main

import (
	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/handler"
	"github.com/dehwyy/Makoto/backend/distributor/logger"
)

func main() {
	// using colored logger
	l := logger.New()

	// looking for 'port' variable in config files
	port, isFound := config.GetOptionByKey("server.distributor")
	if !isFound {
		l.Errorf("port variable was not found")
	}

	// creating new http handler and starting server
	srv := handler.New(port, l)
	if err := srv.ListenAndServe(); err != nil {
		l.Fatalf("failed to start server: %v", err)
	}
}
