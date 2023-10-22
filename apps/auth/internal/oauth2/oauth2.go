package oauth2

import (
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/config"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
)

type OAuth2Provider interface {
	GetToken(token, code string) (*oauth2lib.Token, TokenStatus)
	GetOAuth2UserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error)
	GetProviderName() string
}

type oauth2ProviderResponse struct {
	Id       string // provider id
	Username string
	Email    string
	Picture  string
}

type OAuth2 struct {
	token_repository *repository.TokenRepository
	config           *config.Config
	l                logger.Logger
}

type TokenStatus int
type Endpoint string

const (
	// enum TokenStatus
	Redirect      TokenStatus = iota + 1 // redirect to google's "provide credentials" page
	Success                              // Token was found in db
	InternalError                        // internal error

	// enum OAuth2GoogleEndpoints
	GoogleProfileURL Endpoint = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="
	GithubProfileURL Endpoint = "https://api.github.com/user"
)

func NewOAuth2(token_repo *repository.TokenRepository, config *config.Config, l logger.Logger) *OAuth2 {
	return &OAuth2{
		token_repository: token_repo,
		config:           config,
		l:                l,
	}
}

func (o *OAuth2) GetProviderInstance(provider_name string) OAuth2Provider {
	c := o.config // config
	switch provider_name {
	case "google":
		return newGoogleOAuth2(c.GoogleClientId, c.GoogleClientSecret, c.GoogleRedirectURL, o.token_repository, o.l)
	case "github":
		return newGithubOAuth2(c.GithubClientId, c.GithubClientSecret, c.GithubRedirectUrl, o.token_repository, o.l)
	default:
		return nil
	}
}
