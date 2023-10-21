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

type round_tripper_transport struct {
	response_auth_header string
}

func (tripper *round_tripper_transport) RoundTrip(r *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return response, err
	}

	tripper.response_auth_header = response.Header.Get(AuthorizationHeader)

	return response, nil
}

func AuthorizationMiddleware(url string, l logger.Logger, next http.Handler) http.Handler {
	return WithAuthorizationHeaderMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		transport := &round_tripper_transport{}

		auth_client := auth.NewAuthProtobufClient(url, &http.Client{
			Transport: transport,
		})

		token := r.Header.Get(AuthorizationHeader)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		// make header
		header := make(http.Header)
		header.Set(AuthorizationHeader, token)

		// attach header to context for Request!
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
		ctx = context.WithValue(ctx, auth_userId_key, res.UserId)
		ctx = context.WithValue(ctx, auth_token_key, transport.response_auth_header)

		if err != nil {
			l.Errorf("set ResponseHeader: %v", err)
		}

		// attach context to request
		r = r.WithContext(ctx)

		// serve
		next.ServeHTTP(w, r)
	}))
}

func AuthorizationMiddlewareRead(ctx context.Context) string {
	userId, userId_ok := ctx.Value(auth_userId_key).(string)
	token, token_ok := ctx.Value(auth_token_key).(string)
	if !(userId_ok && token_ok) {
		return ""
	}

	fmt.Printf("Token in Authorization middleware: %v", token)
	twirp.SetHTTPResponseHeader(ctx, AuthorizationHeader, token)

	return userId
}
