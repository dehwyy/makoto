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

func (t *TokenRepository) UpdateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}

func (t *TokenRepository) CreateOrUpdateToken(userId uuid.UUID, username string) (token string, err error) {
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

	// clarify whehter certain token is already linked with user (via user_id)
	// if it is => update instead of creata
	var user_token models.UserToken

	t.db.Model(&user_token).Where("user_id = ?", userId).First(&user_token)
	if user_token.AccessToken != "" {
		return token, t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
			AccessToken:  token,
			RefreshToken: refresh_token,
			Expiry:       exp,
		}).Error
	}

	// simply create
	return token, t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token,
		RefreshToken: refresh_token,
		Expiry:       exp,
		TokenType:    "Bearer",
	}).Error
}

// if authorization method is OAuth2
func (t *TokenRepository) CreateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}
