package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/tools"
	authGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
)

type auth_context_value struct {
	IsAuth       bool
	UserId       string
	AccessToken  string
	RefreshToken string
}

var (
	// Starts with "_" to avoid importable
	_CONTEXT_AUTH_KEY = &auth_context_value{}

	_NO_AUTH_ENDPOINTS = map[string]bool{
		"signUp": true,
		"signIn": true,
	}
)

type graphql_r struct {
	OperationName string `json:"operationName"`
}

func responseWithZeroContext(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// making context with val and put "IsAuth": "False"
	ctx := context.WithValue(r.Context(), _CONTEXT_AUTH_KEY, &auth_context_value{
		// literally it is not neccessary to put <"IsAuth": false> cuz <false> is a default value, but it is just for the clarification
		IsAuth: false,
	})
	r = r.WithContext(ctx)

	next.ServeHTTP(w, r)
}

// return whether isOk (or not)
func checkEndpoint(bodyIO io.ReadCloser) (originBody []byte, isOK bool) {
	// Here we are getting <operationName> to check endpoint

	// getting []byte of <BOdy>
	rawBody, _ := io.ReadAll(bodyIO)

	var op graphql_r
	// getting hitsujouna fields
	json.Unmarshal(rawBody, &op)

	return rawBody, !_NO_AUTH_ENDPOINTS[op.OperationName]
}

func getTokenFromHeader(token string) (string, error) {
	split_token := strings.Split(token, " ")

	// If user doesn't have token in provided Header[key]
	if token == "" || len(split_token) < 2 {
		return "", fmt.Errorf("invalid token: %s", token)
	}

	// taking only second word cuz token is line "Bearer <token>"
	token = split_token[1]

	return token, nil
}

func (m *middleware) Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			originBody, isOk := checkEndpoint(r.Body)

			// actually, we are resetting body to its original state
			r.Body = io.NopCloser(bytes.NewBuffer(originBody))

			// if current endpoints is in <_NO_AUTH_ENDPOINTS> -> skip
			if !isOk {
				responseWithZeroContext(w, r, next)
				return
			}

			auth_token, err := getTokenFromHeader(r.Header.Get("Authorization"))

			// error would appear if header (or potentially token) is not valid
			if err != nil || auth_token == "" {
				responseWithZeroContext(w, r, next)
				return
			}

			// Getting AuthService addr
			authHost, _ := config.GetOptionByKey("docker_services.auth")
			authPort, _ := config.GetOptionByKey("server.auth")
			authAddr := fmt.Sprintf("%s:%s", authHost, authPort)

			// Connecting to it
			g := tools.NewGrpcTools()
			g.CreateConnection(authAddr, m.l)
			defer g.Conn.Close()

			// creating context that would leave in 5 sec
			ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			cl := authGrpc.NewUserClient(g.Conn)
			res, err := cl.ValidateAuth(ctx, &authGrpc.AccessToken{
				AccessToken: auth_token,
			})

			// ! if no error occured => next handler
			if err == nil {
				ctx = context.WithValue(r.Context(), _CONTEXT_AUTH_KEY, &auth_context_value{
					IsAuth:       true, // cuz no error occured
					UserId:       res.UserId,
					AccessToken:  res.AccessToken,
					RefreshToken: res.RefreshToken,
				})

				// and call the next with our new context
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// ! If AccessToken is not valid (expired), error would appear => App will try to recreate Tokens using RefreshToken

			refresh_token, err := getTokenFromHeader(r.Header.Get("RefreshToken"))
			fmt.Println("REFRESH", refresh_token, err)

			// if err is not nil => RefreshToken is not valid (it either doesn't exist or malformed)
			if err != nil || refresh_token == "" {
				responseWithZeroContext(w, r, next)
				return
			}

			ctx, cancel = context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			res, err = cl.RecreateTokensByRefreshToken(ctx, &authGrpc.RefreshToken{
				RefreshToken: refresh_token,
			})

			// if something went wrong OR refreshToken is not valid
			if err != nil {
				responseWithZeroContext(w, r, next)
				return
			}

			ctx = context.WithValue(r.Context(), _CONTEXT_AUTH_KEY, &auth_context_value{
				IsAuth:       true, // cuz no error occured
				UserId:       res.UserId,
				AccessToken:  res.AccessToken,
				RefreshToken: res.RefreshToken,
			})

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ReadAuthContext(ctx context.Context) *auth_context_value {
	// I guess even if read fails, val would be nil, so it satisfies our condition
	val, _ := ctx.Value(_CONTEXT_AUTH_KEY).(*auth_context_value)

	return val
}
