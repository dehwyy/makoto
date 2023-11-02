package middleware

import (
	"context"
	"net/http"

	"github.com/dehwyy/makoto/libs/logger"
)

type onlyAuthorized struct {
	md *withAuthorization
}

func NewMiddleware_OnlyAuthorized(url string, l logger.Logger) *onlyAuthorized {
	return &onlyAuthorized{
		md: NewMiddleware_WithAuthorization(url, l),
	}
}

func (s *onlyAuthorized) Middleware(next http.Handler) http.Handler {
	return s.md.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := s.md.Read(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}))
}

func (s *onlyAuthorized) Read(ctx context.Context) (userId, token string) {
	userId, token, _ = s.md.Read(ctx)
	return userId, token
}
