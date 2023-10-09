package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
	oauth2google "golang.org/x/oauth2/google"
)

type google struct {
	l                logger.Logger
	config           *oauth2lib.Config
	token_repository *repository.TokenRepository
}

type TokenStatus int
type GoogleEndpoint string

const (
	// enum TokenStatus
	Invalid       TokenStatus = iota + 1
	Redirect                  // redirect to google's "provide credentials" page
	Success                   // Token was found in db
	InternalError             // internal error

	// enum OAuth2GoogleEndpoints
	GoogleProfile GoogleEndpoint = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="
)

func NewGoogleOAuth2(clientId, secret, redirectURL string, token_repository *repository.TokenRepository, logger logger.Logger) *google {
	return &google{
		config: &oauth2lib.Config{
			ClientID:     clientId,
			ClientSecret: secret,
			Endpoint:     oauth2google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			RedirectURL:  redirectURL,
		},
		token_repository: token_repository,
		l:                logger,
	}
}

func (g *google) GetToken(access_token *string) (*oauth2lib.Token, TokenStatus) {
	ctx := context.Background()

	if access_token == nil {
		code := g.getCode(ctx) // TODO: should return Redirect
		token := g.createTokenByCode(code)
		// TODO: save to db
		return token, Success
	}

	token_from_db, userId := g.token_repository.GetToken(*access_token)
	oauth2_token := pipes.ToOAuth2TokenFromDbModel(token_from_db)

	// clarify whether Token is not expired
	if token_from_db.Expiry.After(time.Now()) {
		return oauth2_token, Success
	}

	new_token, err := g.config.TokenSource(ctx, oauth2_token).Token() // renew token OR
	if err != nil {
		g.l.Errorf("token renew: %v", err)
		return nil, Redirect
	}

	// TODO: I guess new_token.AccessToken would be different from token_from_db.AccessToken as It's a new token xd
	if new_token.AccessToken == token_from_db.AccessToken {
		return new_token, Success
	}

	err = g.token_repository.SaveToken(*userId, new_token)
	if err != nil {
		g.l.Errorf("token save: %v", err)
		return nil, InternalError
	}

	return new_token, Success
}

func (g *google) DoRequest(endpoint GoogleEndpoint, token *oauth2lib.Token) (*http.Response, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)
	url := string(endpoint) + token.AccessToken

	res, err := cl.Get(url)

	return res, err
}

func (g *google) createTokenByCode(code string) *oauth2lib.Token {
	ctx := context.Background()

	token, err := g.config.Exchange(ctx, code)
	if err != nil {
		g.l.Errorf("token exchange: %v", err)
	}

	return token
}

func (g *google) getCode(ctx context.Context) string {
	// @see https://pkg.go.dev/golang.org/x/oauth2@v0.13.0#GenerateVerifier
	// v := oauth2lib.GenerateVerifier()

	// creating url to Authorize
	url := g.config.AuthCodeURL("state", oauth2lib.AccessTypeOffline /*oauth2lib.S256ChallengeOption(v)*/)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// // TODO
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		g.l.Fatalf("scan stdin: %v", err)
	}

	return code
}
