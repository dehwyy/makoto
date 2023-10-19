package middleware

import (
	"context"
	"net/http"
)

func WithAuthorizationHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get(AuthorizationHeader)
		ctx = context.WithValue(ctx, auth_token_key, token)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func ReadCtxAuthorizationHeader(ctx context.Context) string {
	token, is_ok := ctx.Value(auth_token_key).(string)
	if !is_ok {
		return ""
	}

	return token
}
