package oauth2

import (
	"context"
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
	token_repository repository.TokenRepositoryReadonly
}

func newGoogleOAuth2(clientId, secret, redirectURL string, token_repository repository.TokenRepositoryReadonly, logger logger.Logger) *google {
	return &google{
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

func (g *google) GetToken(access_token, code string) (*oauth2lib.Token, TokenStatus) {
	ctx := context.Background()

	if access_token == "" {
		token := g.createTokenByCode(code)
		return token, Success
	}

	token_from_db, err := g.token_repository.GetToken(access_token)
	if err != nil {
		g.l.Errorf("token get: %v", err)
		return nil, InternalError
	}

	oauth2_token := pipes.ToOAuth2TokenFromDbModel(token_from_db)

	// clarify whether Token is not expired
	if token_from_db.Expiry.After(time.Now()) {
		return oauth2_token, Success
	}

	new_token, err := g.config.TokenSource(ctx, oauth2_token).Token() // renew token OR
	if err != nil {
		g.l.Errorf("token renew: %v", err)
		return nil, Redirect // no reason to provide 2nd arg - user_id
	}

	return new_token, Success
}

func (g *google) GetProviderName() string {
	return "google"
}

func (g *google) GetOAuth2UserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)
	url := string(GoogleProfileURL) + token.AccessToken

	res, err := cl.Get(url)
	if err != nil {
		g.l.Errorf("request: %v", err)
		return nil, err
	}

	var GoogleResponse struct {
		Id      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
		// email:dehwyy@gmail.com
		// given_name:dehwyy
		// id: 103623406957472659690
		// name:dehwyy
		// picture: `https://lh3.googleusercontent.com/a/ACg8ocLE4oqn1c6KC1jgzJB3vL3hhJBDEKxINbHfQmG34Ubrozk=s96-c`
	}

	if err := pipes.Body2Struct(res.Body, &GoogleResponse); err != nil {
		g.l.Errorf("pipes res.body %v", err)
		return nil, err
	}

	return &oauth2ProviderResponse{
		Id:       GoogleResponse.Id,
		Username: GoogleResponse.Name,
		Email:    GoogleResponse.Email,
		Picture:  GoogleResponse.Picture,
	}, nil
}

func (g *google) createTokenByCode(code string) *oauth2lib.Token {
	ctx := context.Background()

	token, err := g.config.Exchange(ctx, code, oauth2lib.AccessTypeOffline)
	if err != nil {
		g.l.Errorf("token exchange: %v", err)
	}

	g.l.Debugf("TOKEN EXCHANGE: %v", token)

	return token
}
