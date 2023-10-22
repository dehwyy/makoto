package oauth2

import (
	"context"
	"time"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
)

type oauth2ProviderResponse struct {
	Id       string // provider id
	Username string
	Email    string
	Picture  string
}

type OAuth2Provider interface {
	GetUserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error)
	GetProviderName() string
}

type OAuth2 struct {
	l                logger.Logger
	config           *oauth2lib.Config
	token_repository repository.TokenRepositoryReadonly
	provider         OAuth2Provider
}

type TokenStatus int

const (
	// enum TokenStatus
	Redirect      TokenStatus = iota + 1 // redirect to "provide credentials" page
	Success                              // Token was found in db
	InternalError                        // internal error
)

// ? constructor is in ./oauth2_creator.go

func (g *OAuth2) GetProviderName() string {
	return g.provider.GetProviderName()
}

func (g *OAuth2) GetUserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
	return g.provider.GetUserByToken(token)
}

func (o *OAuth2) GetToken(access_token, code string) (*oauth2lib.Token, TokenStatus) {
	ctx := context.Background()

	if access_token == "" {
		token := o.createTokenByCode(code)
		return token, Success
	}

	token_from_db, err := o.token_repository.GetToken(access_token)
	if err != nil {
		o.l.Errorf("token get: %v", err)
		return nil, InternalError
	}

	oauth2_token := pipes.ToOAuth2TokenFromDbModel(token_from_db)

	// clarify whether Token is not expired
	if token_from_db.Expiry.After(time.Now()) {
		return oauth2_token, Success
	}

	new_token, err := o.config.TokenSource(ctx, oauth2_token).Token() // renew token OR
	if err != nil {
		o.l.Errorf("token renew: %v", err)
		return nil, Redirect // no reason to provide 2nd arg - user_id
	}

	return new_token, Success
}
func (o *OAuth2) createTokenByCode(code string) *oauth2lib.Token {
	ctx := context.Background()

	token, err := o.config.Exchange(ctx, code, oauth2lib.AccessTypeOffline)
	if err != nil {
		o.l.Errorf("token exchange: %v", err)
	}

	o.l.Debugf("TOKEN EXCHANGE: %v", token)

	return token
}
