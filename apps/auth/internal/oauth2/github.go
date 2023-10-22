package oauth2

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

type github struct {
	l      logger.Logger
	config *oauth2lib.Config
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
		l: logger,
	}
}

func (g *github) GetUserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
	ctx := context.Background()
	cl := g.config.Client(ctx, token)

	req, _ := http.NewRequest("GET", GithubProfileURL, nil)
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
