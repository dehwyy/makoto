package middleware

import "errors"

type ctx_key string

const (
	_AuthorizationHeader               = "Authorization"
	_CtxKeyAuthorizationHeader ctx_key = "CtxKeyAuthorizationHeader"
	_CtxKeyUserId              ctx_key = "CtxKeyUserId"
)

var (
	ErrAuthorizationHeaderNotFound = errors.New("authorization header was not found")
	ErrAuthorizationFailed         = errors.New("authorization failed")
)
