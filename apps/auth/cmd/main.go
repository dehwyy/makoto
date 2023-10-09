package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dehwyy/makoto/apps/auth/internal/db"
	"github.com/dehwyy/makoto/apps/auth/internal/oauth2"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/apps/auth/internal/twirp"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/go-chi/chi/v5"
)

var (
	l   = logger.New()
	cfg = config.New("../../../")
)

func main() {
	l.Debugf("config dsn: %+v", cfg.Databases.Auth)

	r := chi.NewRouter()

	twirp := twirp.NewTwirpServer()

	r.Handle(twirp.PathPrefix(), twirp)

	l.Fatalf("server shutdown, %v", http.ListenAndServe(":"+cfg.Ports.Auth, r))
}

func t() {
	db := db.New(cfg.Databases.Auth, l)
	tr := repository.NewTokenRepository(db, l)
	google_oauth2 := oauth2.NewGoogleOAuth2(cfg.Oauth2.Google.Id, cfg.Oauth2.Google.Secret, cfg.Oauth2.Google.RedirectURL, tr, l)
	token, status := google_oauth2.GetToken(nil)
	if status != oauth2.Success {
		l.Errorf("get token: %v", status)
	}

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
