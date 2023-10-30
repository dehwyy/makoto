package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/twitchtv/twirp"
)

type withAuthorization struct {
	authorizationClientUrl string
	l                      logger.Logger
}

func NewMiddleware_WithAuthorization(authorizationClientUrl string, l logger.Logger) *withAuthorization {
	return &withAuthorization{
		authorizationClientUrl: authorizationClientUrl,
		l:                      l,
	}
}

func (middleware *withAuthorization) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// if token is not provided in header
		token := r.Header.Get(_AuthorizationHeader)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		// make header
		header := make(http.Header)
		header.Set(_AuthorizationHeader, token)

		// attach header to context for request
		ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
		if err != nil {
			fmt.Println("failed to attach header to context in AuthorizationMiddleware")
			next.ServeHTTP(w, r)
			return
		}

		transport := newTwirpClientRoundTripper()
		twirpAuthorizationClient := auth.NewAuthProtobufClient(middleware.authorizationClientUrl, &http.Client{
			Transport: transport,
		})

		res, err := twirpAuthorizationClient.SignIn(ctx, &auth.SignInRequest{
			AuthMethod: &auth.SignInRequest_Empty{
				Empty: &empty.Empty{},
			},
		})
		if err != nil {
			middleware.l.Errorf("failed to call SignIn in AuthorizationMiddleware: %v", err)
			next.ServeHTTP(w, r)
			return
		}

		// set value to ctx
		ctx = context.WithValue(ctx, _CtxKeyUserId, res.UserId)
		ctx = context.WithValue(ctx, _CtxKeyAuthorizationHeader, transport.AuthorizationHeader)

		// attach context to request
		r = r.WithContext(ctx)

		// serve
		next.ServeHTTP(w, r)
	})
}

func (middleware *withAuthorization) Read(ctx context.Context) (userId, token string, err error) {
	userId, userId_ok := ctx.Value(_CtxKeyUserId).(string)
	token, token_ok := ctx.Value(_CtxKeyAuthorizationHeader).(string)
	if !(userId_ok && token_ok) {
		return "", "", ErrAuthorizationFailed
	}

	middleware.l.Infof("userId: %v, token: %v", userId, token)

	return userId, token, nil
}
