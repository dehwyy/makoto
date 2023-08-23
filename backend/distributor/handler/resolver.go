package handler

import (
	"github.com/dehwyy/Makoto/backend/distributor/graphql"
	"github.com/dehwyy/Makoto/backend/distributor/logger"
)

type Resolver struct {
	log logger.AppLogger
}

type mutResolver struct {
	*Resolver
	log logger.AppLogger
}
type queryResolver struct {
	*Resolver
	log logger.AppLogger
}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutResolver{r, r.log}
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r, r.log}
}
