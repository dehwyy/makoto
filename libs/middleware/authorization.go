package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/twitchtv/twirp"
)

func AuthorizationMiddleware(url string, next http.Handler) http.Handler {
	return WithAuthorizationHeaderMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth_client := auth.NewAuthProtobufClient(url, &http.Client{})

		token := WithAuthorizationHeaderMiddlewareRead(ctx)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		// make header
		header := make(http.Header)
		header.Set("Authorization", token)

		// attach header to context
		ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
		if err != nil {
			fmt.Println("failed to attach header to context in AuthorizationMiddleware")
			next.ServeHTTP(w, r)
			return
		}

		res, err := auth_client.SignIn(ctx, &auth.SignInRequest{
			AuthMethod: &auth.SignInRequest_Empty{
				Empty: &empty.Empty{},
			},
		})
		if err != nil {
			fmt.Println("failed to call SignIn in AuthorizationMiddleware")
			next.ServeHTTP(w, r)
			return
		}
		fmt.Printf("response id %v\n", res)

		// set value to ctx
		ctx = context.WithValue(ctx, _AuthorizationKey, res.UserId)

		// attach context to request
		r = r.WithContext(ctx)

		// serve
		next.ServeHTTP(w, r)
	}))
}

func AuthorizationMiddlewareRead(ctx context.Context) string {
	v, isOk := ctx.Value(_AuthorizationKey).(string)
	if !isOk {
		return ""
	}

	return v
}
