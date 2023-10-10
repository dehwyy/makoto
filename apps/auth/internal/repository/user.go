package repository

import (
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateUserPayload struct {
	// UserData
	ID       uuid.UUID
	Id       string
	Email    string
	Name     string
	Picture  string
	Password string

	// UserSpecification
	Provider string
}

type UserRepository struct {
	db *gorm.DB
	l  logger.Logger
}

func NewUserRepository(db *gorm.DB, l logger.Logger) *UserRepository {

	return &UserRepository{
		db: db,
		l:  l,
	}
}

func (u *UserRepository) CreateUser(user_payload CreateUserPayload) error {
	return u.db.Model(&models.UserData{}).Create(&models.UserData{
		ID:         user_payload.ID,
		Username:   user_payload.Name,
		Email:      user_payload.Email,
		CustomId:   user_payload.Id,
		Picture:    user_payload.Picture,
		Provider:   user_payload.Provider,
		Role:       "user",
		ProviderId: user_payload.Id,
		Password:   user_payload.Password,
	}).Error
}
