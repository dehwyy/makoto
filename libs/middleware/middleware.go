package middleware

type MiddlewareCtxKeys int

const (
	// ctx keys
	auth_token_key MiddlewareCtxKeys = iota + 1
	auth_userId_key

	//
	AuthorizationHeader = "Authorization"
)
