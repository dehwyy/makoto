package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/redis/go-redis/v9"
	"github.com/twitchtv/twirp"
)

type withAuthorization struct {
	authorizationClientUrl string
	l                      logger.Logger
	redis                  *redis.Client
}

func NewMiddleware_WithAuthorization(authorizationClientUrl string, rds *redis.Client, l logger.Logger) *withAuthorization {
	return &withAuthorization{
		authorizationClientUrl: authorizationClientUrl,
		l:                      l,
		redis:                  rds,
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

		var userId string
		token = split_token[1]

		// try to get values from redis
		// looks like {"token": "userId"}
		redis_value, err := middleware.redis.Get(ctx, token).Result()

		middleware.l.Infof("redis value: for %s, %s", token, redis_value)

		if err == redis.Nil {
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

			// cache for 10 minutes
			middleware.redis.Set(ctx, token, res.UserId, time.Minute*10)

			userId = res.UserId
			token = res.Token
		} else {
			// if value was found
			userId = redis_value
		}

		// set value to ctx
		ctx = context.WithValue(ctx, _CtxKeyUserId, userId)
		ctx = context.WithValue(ctx, _CtxKeyAuthorizationHeader, token)

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
