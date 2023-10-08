package repository

import (
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/models"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

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

func (t *TokenRepository) GetToken(access_token string) (*models.UserToken, *uuid.UUID) {
	return nil, nil
}

func (t *TokenRepository) SaveToken(userId uuid.UUID, token *oauth2.Token) error {
	return nil
}
