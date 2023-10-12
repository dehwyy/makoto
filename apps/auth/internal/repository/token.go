package repository

import (
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
	db *gorm.DB
	l  logger.Logger
}

func NewTokenRepository(db *gorm.DB, l logger.Logger) *TokenRepository {
	return &TokenRepository{
		db: db,
		l:  l,
	}
}

func (t *TokenRepository) GetToken(access_token string) (*models.UserToken, error) {
	var token models.UserToken

	res := t.db.Model(&models.UserToken{}).Where("access_token = ?", access_token).First(&token)

	return &token, res.Error
}

func (t *TokenRepository) UpdateToken(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}

func (t *TokenRepository) CreateToken(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}
