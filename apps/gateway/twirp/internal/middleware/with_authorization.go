package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
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
		split_token := strings.Split(token, " ")
		if token == "" || len(split_token) != 2 {
			next.ServeHTTP(w, r)
			return
		}

		token = split_token[1]

		twirpAuthorizationClient := auth.NewAuthRPCProtobufClient(middleware.authorizationClientUrl, &http.Client{})

		res, err := twirpAuthorizationClient.SignIn(ctx, &auth.SignInRequest{
			AuthMethod: &auth.SignInRequest_Token{
				Token: token,
			},
		})
		if err != nil {
			middleware.l.Errorf("failed to call SignIn in AuthorizationMiddleware: %v", err)
			next.ServeHTTP(w, r)
			return
		}

		// set value to ctx
		ctx = context.WithValue(ctx, _CtxKeyUserId, res.UserId)
		ctx = context.WithValue(ctx, _CtxKeyAuthorizationHeader, res.Token)

		// attach context to request
		r = r.WithContext(ctx)

		// serve
		next.ServeHTTP(w, r)
	})
}

func (middleware *withAuthorization) Read(ctx context.Context) AuthCredentials {
	userId, userId_ok := ctx.Value(_CtxKeyUserId).(string)
	token, token_ok := ctx.Value(_CtxKeyAuthorizationHeader).(string)
	if !(userId_ok && token_ok) {
		return new_auth_credentials("", "", ErrAuthorizationHeaderNotFound)
	}

	middleware.l.Infof("userId: %v, token: %v", userId, token)
	twirp.SetHTTPResponseHeader(ctx, _AuthorizationHeader, token)
	return new_auth_credentials(userId, token, nil)
}
