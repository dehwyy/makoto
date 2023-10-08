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
	Invalid TokenStatus = iota + 1
	NotYetInDb
	AlreadyInDb

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
		return g.authorizeGoogle(ctx), NotYetInDb
	}

	token_from_db, userId := g.token_repository.GetToken(*access_token)
	oauth2_token := pipes.ToOAuth2TokenFromDbModel(token_from_db)

	// clarify whether Token is not expired
	if token_from_db.Expiry.After(time.Now()) {
		return oauth2_token, AlreadyInDb
	}

	new_token, err := g.config.TokenSource(ctx, oauth2_token).Token() // renew token OR
	if err != nil {
		g.l.Errorf("token renew: %v", err)
		return nil, Invalid
	}

	// TODO: I guess new_token.AccessToken would be different from token_from_db.AccessToken as It's a new token xd
	if new_token.AccessToken == token_from_db.AccessToken {
		return new_token, NotYetInDb
	}

	err = g.token_repository.SaveToken(*userId, new_token)
	if err != nil {
		g.l.Errorf("token save: %v", err)
		return nil, Invalid
	}

	return new_token, NotYetInDb
}

func (g *google) DoRequest(endpoint GoogleEndpoint, token *oauth2lib.Token) (*http.Response, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)
	url := string(endpoint) + token.AccessToken

	res, err := cl.Get(url)

	return res, err
}

func (g *google) authorizeGoogle(ctx context.Context) *oauth2lib.Token {
	// @see https://pkg.go.dev/golang.org/x/oauth2@v0.13.0#GenerateVerifier
	v := oauth2lib.GenerateVerifier()

	// creating url to Authorize
	url := g.config.AuthCodeURL("state", oauth2lib.AccessTypeOffline, oauth2lib.S256ChallengeOption(v))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// TODO
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		g.l.Fatalf("scan stdin: %v", err)
	}

	token, err := g.config.Exchange(ctx, code, oauth2lib.VerifierOption(v))
	if err != nil {
		g.l.Fatalf("token exchange: %v", err)
	}

	return token
}
