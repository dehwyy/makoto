package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dehwyy/Makoto/backend/distributor/config"
	"github.com/dehwyy/Makoto/backend/distributor/tools"
	authGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
)

type auth_context_value struct {
	IsAuth bool
	UserId string
}

var (
	// Starts with "_" to avoid importable
	_CONTEXT_AUTH_KEY = &auth_context_value{}
)

func (m *middleware) Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth_token := r.Header.Get("Authorization")

			// If user doesn't have auth token
			if auth_token == "" {

				// making context with val and put "IsAuth": "False"
				ctx := context.WithValue(r.Context(), _CONTEXT_AUTH_KEY, &auth_context_value{
					IsAuth: false,
					UserId: "",
				})
				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
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
			res, _ := cl.ValidateAuth(ctx, &authGrpc.AccessToken{
				AccessToken: auth_token,
			})

			ctx = context.WithValue(r.Context(), _CONTEXT_AUTH_KEY, &auth_context_value{
				IsAuth: res.IsOk,
				UserId: res.UserId,
			})

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ReadAuthContext(ctx context.Context) (*auth_context_value, bool) {
	val, isOk := ctx.Value(_CONTEXT_AUTH_KEY).(*auth_context_value)
	return val, isOk
}
