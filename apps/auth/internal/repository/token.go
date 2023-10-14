package repository

import (
	"errors"

	"github.com/dehwyy/makoto/apps/auth/internal/utils"
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type TokenRepositoryReadonly interface {
	GetToken(access_token string) (*models.UserToken, error)
}

type TokenRepository struct {
	db  *gorm.DB
	jwt *utils.Jwt
	l   logger.Logger
}

func NewTokenRepository(db *gorm.DB, l logger.Logger, jwt_secret string) *TokenRepository {
	return &TokenRepository{
		db:  db,
		jwt: utils.NewJwt(jwt_secret),
		l:   l,
	}
}

func (t *TokenRepository) GetToken(access_token string) (*models.UserToken, error) {
	var token models.UserToken

	res := t.db.Model(&models.UserToken{}).Where("access_token = ?", access_token).First(&token)

	return &token, res.Error
}

// ! if authorization method is credentials
func (t *TokenRepository) CreateToken(userId uuid.UUID, username string) (token string, err error) {
	jwtPayload := utils.JwtPayload{
		Username: username,
		UserId:   userId.String(),
	}

	token, exp, err_access := t.jwt.NewAccessToken(jwtPayload)
	refresh_token, _, err_refresh := t.jwt.NewRefreshToken(jwtPayload)
	if err_access != nil || err_refresh != nil {
		t.l.Errorf("Failder to create token %v, %v", err_access, err_refresh)
		return "", errors.New("failed to create token")
	}

	return token, t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token,
		RefreshToken: refresh_token,
		Expiry:       exp,
		TokenType:    "Bearer",
	}).Error
}

func (t *TokenRepository) UpdateToken(userId uuid.UUID, username string) (token string, err error) {
	token, exp, err := t.jwt.NewAccessToken(utils.JwtPayload{
		Username: username,
		UserId:   userId.String(),
	})

	if err != nil {
		return "", err
	}

	return token, t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken: token,
		Expiry:      exp,
	}).Error
}

func (t *TokenRepository) ValidateToken(token string) error {
	_, err := t.jwt.ValidateJwtToken(token)
	return err
}

// ! if authorization method is OAuth2
func (t *TokenRepository) UpdateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}

func (t *TokenRepository) CreateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}
