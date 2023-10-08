package main

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/dehwyy/makoto/apps/auth/internal/db"
	"github.com/dehwyy/makoto/apps/auth/internal/oauth2"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/logger"
)

var (
	l   = logger.New()
	cfg = config.New("../../../")
)

func main() {
	// var config = &oauth2lib.Config{
	// 	ClientID:     cfg.Oauth2.Google.Id,
	// 	ClientSecret: cfg.Oauth2.Google.Secret,
	// 	Endpoint:     google.Endpoint,
	// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	// 	RedirectURL:  cfg.Oauth2.Google.RedirectURL,
	// }
	l.Debugf("config: %+v", cfg.Databases.Auth)

	db := db.New(cfg.Databases.Auth, l)

	tr := repository.NewTokenRepository(db, l)

	google_oauth2 := oauth2.NewGoogleOAuth2(cfg.Oauth2.Google.Id, cfg.Oauth2.Google.Secret, cfg.Oauth2.Google.RedirectURL, tr, l)

	token, _ := google_oauth2.GetToken(nil)

	l.Infof("token: %+v", token)

	res, err := google_oauth2.DoRequest(oauth2.GoogleProfile, token)
	if err != nil {
		l.Fatalf("request: %v", err)
		return
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		l.Fatalf("res body copy %v", err)
		return
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		l.Fatalf("res body unmarshal %v", err)
		return
	}

	l.Debugf("ServicesResponse: %+v", GoogleUserRes)
}
