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
	UserNotFound      = errors.New("User not found")
	UserWrongPassword = errors.New("Wrong password")
	UserAlreadyExists = errors.New("User already exists")
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
	UserId   uuid.UUID
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

// GetUserById retrieves a user by their ID or custom ID.
//
// It takes in a GetUserPayload object which contains the user ID or custom ID.
// The function returns a user object of type UserData and an error object.
func (u *UserRepository) GetUserById(user_payload GetUserPayload) (user *models.UserData, erorr error) {
	if user_payload.Id != nil {
		return user, u.db.Model(&models.UserData{}).Where("id = ?", *user_payload.Id).First(&user).Error
	}

	return user, u.db.Model(&models.UserData{}).Where("custom_id = ?", user_payload.CustomId).First(&user).Error
}

// GetUserByUsername retrieves a user from the UserRepository based on the given username.
//
// Parameters:
// - username: a string representing the username of the user to retrieve.
//
// Returns:
// - user: a pointer to a UserData struct representing the retrieved user.
// - error: an error if there was an issue retrieving the user.
func (u *UserRepository) GetUserByUsername(username string) (user *models.UserData, erorr error) {
	return user, u.db.Model(&models.UserData{}).Where("username = ?", username).First(&user).Error
}

// GetUserByProviderId retrieves a user by their provider ID.
//
// provider_id: The provider ID of the user.
// Returns:
//
//	user: The user with the specified provider ID.
//	error: Any error that occurred during retrieval.
func (u *UserRepository) GetUserByProviderId(provider_id string) (user *models.UserData, erorr error) {
	return user, u.db.Model(&models.UserData{}).Where("provider_id = ?", provider_id).First(&user).Error
}

// CreateUser creates a new user in the UserRepository.
//
// It takes a CreateUserPayload as a parameter and returns an error.
func (u *UserRepository) CreateUser(user_payload CreateUserPayload) error {
	username := user_payload.Username

	var found_user_by_username *models.UserData // clarify if there is a user with this username
	err := u.db.Model(&models.UserData{}).Where("username = ?", username).First(&found_user_by_username).Error

	// if internal error => return it
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// if err == gorm.ErrRecordNotFound => Ok => won't change username
	// if user was found ( err == nil ) => Error would appear, have to change username (as it must be unique) (non-local provider)
	if err == nil {
		// if provider == local and non-unique username -> throw error
		if user_payload.Provider == ProviderLocal {
			return UserAlreadyExists
		}
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

// ValidateUser validates a user based on the given user payload.
//
// It takes a parameter of type ValidateUserPayload which represents the user payload
// containing email, username, and user ID.
//
// It returns a pointer to a UUID and an error. The UUID represents the user's ID,
// while the error indicates any validation errors that occurred during the process.
func (u *UserRepository) ValidateUser(user_payload ValidateUserPayload) (id *uuid.UUID, username string, err error) {
	var user_data models.UserData

	if user_payload.Email != "" {
		u.db.Model(&models.UserData{}).Where("email = ?", user_payload.Email).First(&user_data)
	} else if user_payload.Username != "" {
		u.db.Model(&models.UserData{}).Where("username = ?", user_payload.Username).First(&user_data)
	} else {
		u.db.Model(&models.UserData{}).Where("id = ?", user_payload.UserId).First(&user_data)
	}

	// if username is equal to "" => user wasn't found as "" is a default string's value
	if user_data.Username == "" {
		return nil, "", UserNotFound
	}

	is_valid := u.hasher.Compare(user_payload.Password, user_data.Password)
	if !is_valid {
		return nil, "", UserWrongPassword
	}

	return &user_data.ID, user_data.Username, nil

}

// VerifyUserEmail verifies the email of a user.
//
// Parameters:
// - user_id: The UUID of the user.
//
// Returns:
// - error: An error if the verification fails.
func (u *UserRepository) VerifyUserEmail(user_id uuid.UUID) error {
	return u.db.Model(&models.UserData{}).Where("id = ?", user_id).Update("is_verified", true).Error
}

// UpdateUserPassword updates the password of a user in the UserRepository.
//
// Parameters:
// - user_id: The ID of the user.
// - new_password: The new password for the user.
//
// Return type:
// - error: An error if the update operation fails.
func (u *UserRepository) UpdateUserPassword(user_id uuid.UUID, new_password string) error {
	return u.db.Model(&models.UserData{}).Where("id = ?", user_id).Update("password", new_password).Error
}

// DeleteUser deletes a user from the repository.
//
// Parameters:
// - user_id: The ID of the user to be deleted.
//
// Returns:
// - error: An error if the deletion operation fails.
func (u *UserRepository) DeleteUser(user_id uuid.UUID) error {
	return u.db.Model(&models.UserData{}).Where("id = ?", user_id).Delete(&models.UserData{}).Error
}
