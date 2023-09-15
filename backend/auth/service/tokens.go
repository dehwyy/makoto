package service

import (
	"github.com/dehwyy/Makoto/backend/auth/db/models"
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/dehwyy/Makoto/backend/auth/tools"
	"gorm.io/gorm"
)

type jwtPayload = tools.JwtPayload

type jwtHandler interface {
	// (payload jwtPayload) => (token: string, error: error)
	NewRefreshToken(jwtPayload) (string, error)
	// same as previous
	NewAccessToken(jwtPayload) (string, error)
	// (token: string) => error (if valid => nil)
	ValidateJwtToken(string) error
}

type TokenService struct {
	jwt jwtHandler
	db  *gorm.DB
	l   logger.AppLogger
}

func NewTokenService(jwt jwtHandler, db *gorm.DB, l logger.AppLogger) *TokenService {
	return &TokenService{
		jwt: jwt,
		db:  db,
		l:   l,
	}
}

func (t *TokenService) newJwtPayload(username, userId string) jwtPayload {
	return jwtPayload{
		Username: username,
		UserId:   userId,
	}
}

// Signing access and refresh tokens; The refresh one would be saved to db; Returns (ACCESS_TOKEN, REFRESH_TOKEN).
func (t *TokenService) SignTokensAndSave(username, userId string) (string, string) {
	payload := t.newJwtPayload(username, userId)

	refresh_token, err := t.jwt.NewRefreshToken(payload)
	if err != nil {
		t.l.Errorf("Error creating refresh token: %v", err)
	}

	// saving refresh token to db
	t.db.Create(&models.Token{
		Token:  refresh_token,
		UserId: userId,
	})
	t.l.Infof("Generated refresh token: %v", refresh_token)

	access_token, err := t.jwt.NewAccessToken(payload)
	if err != nil {
		t.l.Errorf("Error creating access token: %v", err)
	}
	t.l.Infof("Generated access token: %v", access_token)

	return access_token, refresh_token
}
