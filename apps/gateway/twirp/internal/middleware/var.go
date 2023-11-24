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

type AuthCredentials interface {
	GetUserId() (string, error)
	GetToken() (string, error)
	GetError() error
}

type AuthCredentialsGranted interface {
	UserId() string
	Token() string
}

type auth_credentials struct {
	userId string
	token  string
	err    error
}

// non-granted
func new_auth_credentials(userId, token string, err error) AuthCredentials {
	return &auth_credentials{
		userId: userId,
		token:  token,
		err:    err,
	}
}

func (a *auth_credentials) GetUserId() (string, error) {
	return a.userId, a.err
}

func (a *auth_credentials) GetToken() (string, error) {
	return a.token, a.err
}

func (a *auth_credentials) GetError() error {
	return a.err
}

// granted
func new_auth_credentials_granted(userId, token string) AuthCredentialsGranted {
	return &auth_credentials{
		userId: userId,
		token:  token,
	}
}

func (a *auth_credentials) UserId() string {
	return a.userId
}

func (a *auth_credentials) Token() string {
	return a.token
}
