package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	ds "github.com/ravener/discord-oauth2"
	oauth2lib "golang.org/x/oauth2"
)

type discord struct {
	config *oauth2lib.Config
	l      logger.Logger
}

func newDiscordOAuth2(clientId, clientSecret, redirectURL string, token_repository repository.TokenRepositoryReadonly, logger logger.Logger) *discord {
	return &discord{
		config: &oauth2lib.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint:     ds.Endpoint,
			RedirectURL:  redirectURL,
			Scopes:       []string{ds.ScopeEmail, ds.ScopeActivitiesRead},
		},
		l: logger,
	}
}

func (d *discord) GetProviderName() string {
	return "discord"
}

func (d *discord) GetUserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
	ctx := context.Background()
	cl := d.config.Client(ctx, token)

	req, _ := http.NewRequest("GET", DiscordProfileURL, nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := cl.Do(req)
	if err != nil {
		d.l.Errorf("request: %v", err)
		return nil, err
	}

	var DiscordResponse struct {
		User struct {
			Id       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Avatar   string `json:"avatar"`
		} `json:"user"`
	}

	if err := pipes.Body2Struct(res.Body, &DiscordResponse); err != nil {
		d.l.Errorf("pipes res.body %v", err)
		return nil, err
	}

	d.l.Debugf("DiscordResponse: %v", DiscordResponse)
	avatar := DiscordResponse.User.Avatar
	if avatar != "" {
		avatar = fmt.Sprintf("https://cdn.discordapp.com/avatars/user_id/%s.png", avatar)
	}

	return &oauth2ProviderResponse{
		Id:       DiscordResponse.User.Id,
		Username: DiscordResponse.User.Username,
		Email:    DiscordResponse.User.Email,
		Picture:  avatar,
	}, nil
}
