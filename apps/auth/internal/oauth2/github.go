package oauth2

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

type github struct {
	l                logger.Logger
	config           *oauth2lib.Config
	token_repository repository.TokenRepositoryReadonly
}

func newGithubOAuth2(clientId, secret, redirectURL string, token_repository repository.TokenRepositoryReadonly, logger logger.Logger) *github {
	return &github{
		config: &oauth2lib.Config{
			ClientID:     clientId,
			ClientSecret: secret,
			Endpoint:     oauth2github.Endpoint,
			Scopes:       []string{},
			RedirectURL:  redirectURL,
		},
		// ! READONLY Repository
		token_repository: token_repository,
		l:                logger,
	}
}

func (g *github) GetToken(token, code string) (*oauth2lib.Token, TokenStatus) {
	ctx := context.Background()

	if token == "" {
		token := g.createTokenByCode(code)
		return token, Success
	}

	token_from_db, err := g.token_repository.GetToken(token)
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

func (g *github) GetOAuth2UserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)

	req, _ := http.NewRequest("GET", string(GithubProfileURL), nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	req.Header.Add("Accept", "application/json")

	res, err := cl.Do(req)
	if err != nil {
		g.l.Errorf("request: %v", err)
		return nil, err
	}

	var GithubResponse struct {
		Id      int    `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"login"`
		Picture string `json:"avatar_url"`
		// email:dehwyy@gmail.com
		// given_name:dehwyy
		// id: 103623406957472659690
		// name:dehwyy
		// picture: `https://lh3.googleusercontent.com/a/ACg8ocLE4oqn1c6KC1jgzJB3vL3hhJBDEKxINbHfQmG34Ubrozk=s96-c`
	}

	if err := pipes.Body2Struct(res.Body, &GithubResponse); err != nil {
		g.l.Errorf("pipes res.body %v", err)
		return nil, err
	}

	g.l.Debugf("GithubResponse: %v", GithubResponse)

	return &oauth2ProviderResponse{
		Id:       strconv.Itoa(GithubResponse.Id),
		Username: GithubResponse.Name,
		Email:    GithubResponse.Email,
		Picture:  GithubResponse.Picture,
	}, nil
}
func (g *github) GetProviderName() string {
	return "github"
}

func (g *github) createTokenByCode(code string) *oauth2lib.Token {
	ctx := context.Background()

	token, err := g.config.Exchange(ctx, code, oauth2lib.AccessTypeOffline)
	if err != nil {
		g.l.Errorf("token exchange: %v", err)
	}

	g.l.Debugf("TOKEN EXCHANGE: %v", token)

	return token
}
