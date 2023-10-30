package main

import (
	"fmt"
	"net/http"

	twirp "github.com/dehwyy/makoto/apps/gateway/twirp/internal/impl"
	"github.com/dehwyy/makoto/apps/gateway/twirp/internal/middleware"
	makoto_config "github.com/dehwyy/makoto/libs/config"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	log := logger.New()
	// config := makoto_config.New()
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	auth_port := fmt.Sprintf("localhost:%v", makoto_config.PortAuth)

	//  middleware that reads the `Authorization` header (as twirp doesn't give access to it directly)
	md_with_authorization_header := middleware.NewMiddleware_WithAuthorizationHeader()
	md_with_authorization := middleware.NewMiddleware_WithAuthorization(auth_port, log)

	// services
	authorization_service := twirp.NewAuthorizationService(twirp.TwirpAuthorizationService{
		ReadHeader: md_with_authorization_header.Read,
	})
	hashmap_service := twirp.NewHashmapService(twirp.TwirpHashmapService{
		ReadAuthorizationData: md_with_authorization.Read,
	})

	// mount
	r.Mount("/authorization", md_with_authorization_header.Middleware(authorization_service))
	r.Mount("/hashmap", md_with_authorization.Middleware(hashmap_service))

	port := ":9000"
	log.Infof("Gateway server started on port %s", port)

	http.ListenAndServe(port, r)
}
