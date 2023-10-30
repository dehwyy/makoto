package middleware

import (
	"context"
	"net/http"
)

type withAuthorizationHeader struct{}

func NewMiddleware_WithAuthorizationHeader() *withAuthorizationHeader {
	return &withAuthorizationHeader{}
}

func (s *withAuthorizationHeader) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get request context
		ctx := r.Context()

		// get token from `Authorization` header
		token := r.Header.Get(_AuthorizationHeader)

		// set Token in context
		ctx = context.WithValue(ctx, _CtxKeyAuthorizationHeader, token)

		// update request
		r = r.WithContext(ctx)

		// call next fn
		next.ServeHTTP(w, r)
	})
}

func (s *withAuthorizationHeader) Read(ctx context.Context) (string, error) {
	// try to get token from ctx
	token, is_ok := ctx.Value(_CtxKeyAuthorizationHeader).(string)

	// if nothing was found -> return error
	if !is_ok {
		return "", ErrAuthorizationHeaderNotFound
	}

	return token, nil
}
