package middleware

import (
	"context"
	"net/http"
)

const (
	_AuthorizationKey MiddlewareKeys = iota + 1
)

func WithAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := r.Header.Get("Authorization")
		ctx = context.WithValue(ctx, _AuthorizationKey, v)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func WithAuthorizationMiddlewareRead(ctx context.Context) string {
	v, isOk := ctx.Value(_AuthorizationKey).(string)
	if !isOk {
		return ""
	}

	return v
}
