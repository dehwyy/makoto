package main

import (
	"net/http"

	"github.com/dehwyy/makoto/apps/auth/internal/twirp"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/database"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var (
	l   = logger.New()
	cfg = config.New("../../../")
)

func main() {
	db := database.New(cfg.Databases.Auth, l)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	twirp := twirp.NewTwirpServer(db, cfg, l)
	r.Mount(twirp.PathPrefix(), middleware.WithAuthorizationMiddleware(twirp))

	l.Infof("Server started on port %v", cfg.Ports.Auth)
	l.Fatalf("server shutdown, %v", http.ListenAndServe(":"+cfg.Ports.Auth, r))
}
