package middleware

type MiddlewareKeys int

const (
	_AuthorizationKey       MiddlewareKeys = iota + 1
	_AuthorizationHeaderKey MiddlewareKeys = 2
)
