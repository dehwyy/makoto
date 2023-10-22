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
			Scopes:       []string{ds.ScopeEmail, ds.ScopeIdentify},
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
		Id         string `json:"id"`
		GlobalName string `json:"global_name"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Avatar     string `json:"avatar"`
	}

	if err := pipes.Body2Struct(res.Body, &DiscordResponse); err != nil {
		d.l.Errorf("pipes res.body %v", err)
		return nil, err
	}
	d.l.Debugf("DiscordResponseUser %v", DiscordResponse)

	// @see https://discord.com/developers/docs/reference#image-formatting
	avatar := DiscordResponse.Avatar
	if avatar != "" {
		avatar = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s?size=1024", DiscordResponse.Id, avatar)
	}

	//
	username := DiscordResponse.Username
	if DiscordResponse.GlobalName != "" {
		username = DiscordResponse.GlobalName
	}

	return &oauth2ProviderResponse{
		Id:       DiscordResponse.Id,
		Username: username,
		Email:    DiscordResponse.Email,
		Picture:  avatar,
	}, nil
}
