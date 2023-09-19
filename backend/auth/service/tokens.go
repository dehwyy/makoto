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
	ValidateJwtToken(string) (*jwtPayload, error)
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

func (t *TokenService) schema() *gorm.DB {
	return t.db.Model(&models.Token{})
}

func (t *TokenService) newJwtPayload(username, userId string) jwtPayload {
	return jwtPayload{
		Username: username,
		UserId:   userId,
	}
}

func (t *TokenService) signTokens(payload jwtPayload) (string, string) {
	// TODO: remove logs

	refresh_token, err := t.jwt.NewRefreshToken(payload)
	if err != nil {
		t.l.Errorf("Error creating refresh token: %v", err)
	}
	t.l.Infof("Generated refresh token: %v", refresh_token)

	access_token, err := t.jwt.NewAccessToken(payload)
	if err != nil {
		t.l.Errorf("Error creating access token: %v", err)
	}
	t.l.Infof("Generated access token: %v", access_token)

	return access_token, refresh_token
}

// Signing access and refresh tokens; The refresh one would be saved to db; Returns (ACCESS_TOKEN, REFRESH_TOKEN).
func (t *TokenService) SignTokensAndCreate(username, userId string) (string, string) {
	payload := t.newJwtPayload(username, userId)

	access_token, refresh_token := t.signTokens(payload)

	// saving refresh token to db
	t.schema().Create(&models.Token{
		Token:  refresh_token,
		UserId: userId,
	})

	return access_token, refresh_token
}

func (t *TokenService) SignTokensAndUpdate(username, userId string) (string, string) {
	payload := t.newJwtPayload(username, userId)

	access_token, refresh_token := t.signTokens(payload)

	// updating refresh token
	t.schema().Where("token = ?", refresh_token).Update("token", refresh_token)

	return access_token, refresh_token
}

// returns UserId | Username | isValid
func (t *TokenService) ValidateToken(token string) (string, string, bool) {
	claims, err := t.jwt.ValidateJwtToken(token)
	if err != nil {
		t.l.Errorf("Error validating token: %v", err)
	}

	return claims.UserId, claims.Username, err == nil
}

func (t *TokenService) RemoveToken(userId string) error {
	res := t.schema().Delete(&models.Token{}, "user_id = ?", userId)

	return res.Error
}
