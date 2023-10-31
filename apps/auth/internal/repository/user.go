package repository

import (
	"errors"
	"fmt"

	"github.com/dehwyy/makoto/apps/auth/internal/utils"
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ProviderGoogle models.AuthProvider = models.ProviderGoogle
	ProviderLocal                      = models.ProviderLocal
)

var (
	USER_NOT_FOUND      = errors.New("User not found")
	USER_WRONG_PASSWORD = errors.New("Wrong password")
)

type GetUserPayload struct {
	Id       *uuid.UUID
	CustomId string
}

type CreateUserPayload struct {
	// UserData
	ID       uuid.UUID
	Id       string
	Email    string
	Username string
	Picture  string
	Password string

	// UserSpecification
	Provider models.AuthProvider
}

type ValidateUserPayload struct {
	Username string
	Email    string
	Password string
}

type UserRepository struct {
	db     *gorm.DB
	hasher utils.Hasher
	l      logger.Logger
}

func NewUserRepository(db *gorm.DB, l logger.Logger) *UserRepository {

	return &UserRepository{
		db: db,
		l:  l,
	}
}

func (u *UserRepository) GetUserById(user_payload GetUserPayload) (user *models.UserData, erorr error) {
	if user_payload.Id != nil {
		return user, u.db.Model(&models.UserData{}).Where("id = ?", *user_payload.Id).First(&user).Error
	}

	return user, u.db.Model(&models.UserData{}).Where("custom_id = ?", user_payload.CustomId).First(&user).Error
}

func (u *UserRepository) GetUserByProviderId(provider_id string) (user *models.UserData, erorr error) {
	return user, u.db.Model(&models.UserData{}).Where("provider_id = ?", provider_id).First(&user).Error
}

func (u *UserRepository) CreateUser(user_payload CreateUserPayload) error {
	username := user_payload.Username

	var found_user_by_username *models.UserData // clarify if there is a user with this username
	err := u.db.Model(&models.UserData{}).Where("username = ?", username).First(&found_user_by_username).Error

	// if internal error => return it
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// if gorm.ErrRecordNotFound => Ok => won't change username
	// if user was found => Error would appear, but we have to change username (as it must be unique)
	if err == nil {
		username = fmt.Sprintf("%s_%s", username, string(user_payload.Provider)) // create new nickname. example: `dehwyy_github` or `dehwyy_google`
	}

	hashed_password, err := u.hasher.Hash(user_payload.Password)
	if err != nil {
		return err
	}

	return u.db.Model(&models.UserData{}).Create(&models.UserData{
		ID:         user_payload.ID,
		Username:   username,
		Email:      user_payload.Email,
		CustomId:   user_payload.Id,
		Picture:    user_payload.Picture,
		Provider:   user_payload.Provider,
		Role:       "user",
		ProviderId: user_payload.Id,
		Password:   hashed_password,
	}).Error
}

func (u *UserRepository) ValidateUser(user_payload ValidateUserPayload) (id *uuid.UUID, err error) {
	var user_data models.UserData

	if user_payload.Email != "" {
		u.db.Model(&models.UserData{}).Where("email = ?", user_payload.Email).First(&user_data)
	} else {
		u.db.Model(&models.UserData{}).Where("username = ?", user_payload.Username).First(&user_data)
	}

	// if username is equal to "" => user wasn't found as "" is a default string's value
	if user_data.Username == "" {
		return nil, USER_NOT_FOUND
	}

	is_valid := u.hasher.Compare(user_payload.Password, user_data.Password)
	if !is_valid {
		return nil, USER_WRONG_PASSWORD
	}

	return &user_data.ID, nil

}
