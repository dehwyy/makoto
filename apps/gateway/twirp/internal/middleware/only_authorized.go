package middleware

import (
	"context"
	"net/http"

	"github.com/dehwyy/makoto/libs/logger"
	"github.com/redis/go-redis/v9"
)

type onlyAuthorized struct {
	md *withAuthorization
}

func NewMiddleware_OnlyAuthorized(url string, rds *redis.Client, l logger.Logger) *onlyAuthorized {
	return &onlyAuthorized{
		md: NewMiddleware_WithAuthorization(url, rds, l),
	}
}

func (s *onlyAuthorized) Middleware(next http.Handler) http.Handler {
	return s.md.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := s.md.Read(r.Context()).GetError()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}))
}

func (s *onlyAuthorized) Read(ctx context.Context) AuthCredentialsGranted {
	token, _ := s.md.Read(ctx).GetToken()
	userId, _ := s.md.Read(ctx).GetUserId()

	return new_auth_credentials_granted(userId, token)
}
