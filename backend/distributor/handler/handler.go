package handler

import (
	"net/http"

	graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dehwyy/Makoto/backend/distributor/graphql"
	"github.com/dehwyy/Makoto/backend/distributor/logger"
	"github.com/dehwyy/Makoto/backend/distributor/middleware"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	srv *http.Server
}

func New(port string, l logger.AppLogger) *handler {
	graphQLServer := createGraphQLServer(l)
	router := chi.NewRouter()

	// Add middleware
	m := middleware.New(l)
	router.Use(m.Auth())

	// Initialize routes for GraphQL
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graphQLServer)

	l.Infof("Connect to graphql playground on http://localhost:%s", port)

	return &handler{
		srv: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
	}
}

func (h *handler) ListenAndServe() error {
	return h.srv.ListenAndServe()
}

func createGraphQLServer(log logger.AppLogger) *graphqlHandler.Server {
	config := graphql.Config{
		Resolvers: &Resolver{
			log: log,
		},
	}

	return graphqlHandler.NewDefaultServer(graphql.NewExecutableSchema(config))
}
