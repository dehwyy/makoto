package oauth2

import (
	"context"

	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/libs/logger"
	oauth2lib "golang.org/x/oauth2"
	oauth2google "golang.org/x/oauth2/google"
)

type google struct {
	l      logger.Logger
	config *oauth2lib.Config
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
		l: logger,
	}
}

func (g *google) GetProviderName() string {
	return "google"
}

func (g *google) GetUserByToken(token *oauth2lib.Token) (*oauth2ProviderResponse, error) {
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
