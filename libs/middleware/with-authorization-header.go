package middleware

import (
	"context"
	"net/http"
)

func WithAuthorizationHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := r.Header.Get("Authorization")
		ctx = context.WithValue(ctx, _AuthorizationHeaderKey, v)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func WithAuthorizationHeaderMiddlewareRead(ctx context.Context) string {
	v, isOk := ctx.Value(_AuthorizationHeaderKey).(string)
	if !isOk {
		return ""
	}

	return v
}
