package oauth2

import (
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/config"
	"github.com/dehwyy/makoto/libs/logger"
	"golang.org/x/oauth2"
)

type OAuth2Creator struct {
	token_repository *repository.TokenRepository
	config           *config.Config
	l                logger.Logger
}

func NewOAuth2Creator(token_repo *repository.TokenRepository, config *config.Config, l logger.Logger) *OAuth2Creator {
	return &OAuth2Creator{
		token_repository: token_repo,
		config:           config,
		l:                l,
	}
}

func (g *OAuth2Creator) NewOAuth2(provider_name string) *OAuth2 {
	var provider OAuth2Provider
	var oauth2config *oauth2.Config

	switch provider_name {
	case "google":
		google := newGoogleOAuth2(g.config.GoogleClientId, g.config.GoogleClientSecret, g.config.GoogleRedirectURL, g.token_repository, g.l)
		provider = google
		oauth2config = google.config
	case "github":
		github := newGithubOAuth2(g.config.GithubClientId, g.config.GithubClientSecret, g.config.GithubRedirectURL, g.token_repository, g.l)
		provider = github
		oauth2config = github.config
	case "discord":
		discord := newDiscordOAuth2(g.config.DiscordClientId, g.config.DiscordClientSecret, g.config.DiscordRedirectURL, g.token_repository, g.l)
		provider = discord
		oauth2config = discord.config
	default:
		g.l.Errorf("unknown provider: %s", provider_name)
		return nil
	}

	return &OAuth2{
		l:                g.l,
		config:           oauth2config,
		token_repository: g.token_repository,
		provider:         provider,
	}
}
