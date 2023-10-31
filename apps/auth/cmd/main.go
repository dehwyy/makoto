package main

import (
	"net/http"
	"strings"

	"github.com/dehwyy/makoto/apps/auth/internal/twirp"
	makoto_config "github.com/dehwyy/makoto/libs/config"
	"github.com/dehwyy/makoto/libs/database"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	log := logger.New()
	config := makoto_config.New()
	db := database.New(config.DatabaseDsn, log)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	twirp := twirp.NewTwirpServer(db, config, log)

	r.Mount(twirp.PathPrefix(), twirp)

	port := ":" + strings.Split(config.AuthUrl, ":")[1]

	log.Infof("Server started on port %v", port)
	log.Fatalf("server shutdown, %v", http.ListenAndServe(port, r))
}
