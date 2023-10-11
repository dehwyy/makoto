package oauth2

import (
	"context"
	"net/http"
	"time"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	oauth2lib "golang.org/x/oauth2"
	oauth2google "golang.org/x/oauth2/google"
)

type Google struct {
	l                logger.Logger
	config           *oauth2lib.Config
	token_repository repository.TokenRepositoryReadonly
}

type TokenStatus int
type GoogleEndpoint string

const (
	// enum TokenStatus
	Redirect      TokenStatus = iota + 1 // redirect to google's "provide credentials" page
	Success                              // Token was found in db
	InternalError                        // internal error

	// enum OAuth2GoogleEndpoints
	GoogleProfile GoogleEndpoint = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="
)

func NewGoogleOAuth2(clientId, secret, redirectURL string, token_repository repository.TokenRepositoryReadonly, logger logger.Logger) *Google {
	return &Google{
		config: &oauth2lib.Config{
			ClientID:     clientId,
			ClientSecret: secret,
			Endpoint:     oauth2google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			RedirectURL:  redirectURL,
		},
		// ! READONLY Repository
		token_repository: token_repository,
		l:                logger,
	}
}

func (g *Google) GetToken(access_token, code string) (*oauth2lib.Token, *uuid.UUID, TokenStatus) {
	ctx := context.Background()

	if access_token == "" {
		token := g.createTokenByCode(code)
		return token, nil, Success
	}

	token_from_db, err := g.token_repository.GetToken(access_token)
	if err != nil {
		g.l.Errorf("token get: %v", err)
		return nil, nil, InternalError
	}

	oauth2_token := pipes.ToOAuth2TokenFromDbModel(token_from_db)

	// clarify whether Token is not expired
	if token_from_db.Expiry.After(time.Now()) {
		return oauth2_token, &token_from_db.UserId, Success
	}

	new_token, err := g.config.TokenSource(ctx, oauth2_token).Token() // renew token OR
	if err != nil {
		g.l.Errorf("token renew: %v", err)
		return nil, nil, Redirect // no reason to provide 2nd arg - user_id
	}

	return new_token, &token_from_db.UserId, Success
}

func (g *Google) DoRequest(endpoint GoogleEndpoint, token *oauth2lib.Token) (*http.Response, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)
	url := string(endpoint) + token.AccessToken

	res, err := cl.Get(url)

	return res, err
}

func (g *Google) createTokenByCode(code string) *oauth2lib.Token {
	ctx := context.Background()

	token, err := g.config.Exchange(ctx, code, oauth2lib.AccessTypeOffline)
	if err != nil {
		g.l.Errorf("token exchange: %v", err)
	}

	g.l.Debugf("TOKEN EXCHANGE: %v", token)

	return token
}
