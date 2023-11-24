package repository

import (
	"errors"
	default_logger "log"
	"os"
	"testing"
	"time"

	"github.com/dehwyy/makoto/libs/database"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

var (
	_log     = logger.New()
	_jwt_key = "secret"
	_db_dsn  = "host=localhost user=postgres password=postgres dbname=makoto port=5432 sslmode=disable"
	_db      = database.New(_db_dsn, _log).Session(&gorm.Session{
		Logger: gorm_logger.New(default_logger.New(os.Stdout, "\r\n", default_logger.LstdFlags), gorm_logger.Config{
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})

	_token_repo = NewTokenRepository(_db, _log, _jwt_key)
	_user_repo  = NewUserRepository(_db, _log)

	_state = &_State{
		name:      "dehwyy_tester",
		user_uuid: uuid.MustParse("193c2043-4058-4323-bcfd-14c356b2a367"),
	}
)

type _State struct {
	name      string
	user_uuid uuid.UUID
	token     string

	token_refresh string
	token_expiry  time.Time
	token_type    string

	old_password_hash string
}

func Test_CreateUser(t *testing.T) {
	_user_repo.DeleteUser(_state.user_uuid)

	_user_repo.CreateUser(CreateUserPayload{
		ID:       _state.user_uuid,
		Id:       _state.user_uuid.String(),
		Picture:  "",
		Provider: "local",
		Username: _state.name,
		Email:    _state.name + "@example.com",
		Password: "password",
	})
}

func Test_CreateToken(t *testing.T) {
	token, err := _token_repo.CreateToken(_state.user_uuid, _state.name)

	// at least 8 bytes (or even more)
	assert.GreaterOrEqual(t, len(token), 16)
	// no error
	assert.Nil(t, err)

	_state.token = token
}

func Test_GetToken(t *testing.T) {
	token, err := _token_repo.GetToken(_state.token)

	// no error
	assert.Nil(t, err)
	// token is not empty
	assert.NotEmpty(t, token)

	// every field is not empty
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
	assert.NotEmpty(t, token.Expiry)
	assert.NotEmpty(t, token.TokenType)
	assert.NotEmpty(t, token.UserId)

	_state.token_expiry = token.Expiry
	_state.token_refresh = token.RefreshToken
	_state.token_type = token.TokenType
}

func Test_UpdateToken(t *testing.T) {
	token, err := _token_repo.UpdateToken(_state.user_uuid)

	// no error
	assert.Nil(t, err)
	// token is not empty
	assert.NotEmpty(t, token)

	// should be new : it should be renewed
	assert.NotEqual(t, _state.token, token)

	// token from db
	new_token, err := _token_repo.GetToken(token)

	// no error
	assert.Nil(t, err)
	// token is not empty
	assert.NotEmpty(t, new_token)

	assert.Equal(t, token, new_token.AccessToken)
	assert.Equal(t, _state.token_refresh, new_token.RefreshToken)
	assert.Equal(t, _state.token_type, new_token.TokenType)
	assert.Equal(t, _state.user_uuid, new_token.UserId)

	// should be renewed
	assert.NotEqual(t, _state.token_expiry, new_token.Expiry)

	_state.token = token
	_state.token_expiry = new_token.Expiry
}

func Test_ValidateToken(t *testing.T) {
	err := _token_repo.ValidateToken(_state.token)

	// no error
	assert.Nil(t, err)
}

func Test_DeleteToken(t *testing.T) {
	err := _token_repo.DeleteToken(_state.user_uuid)

	// no error
	assert.Nil(t, err)

	token, err := _token_repo.GetToken(_state.token)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		assert.Nil(t, token)
	} else {
		assert.Fail(t, err.Error())
	}
}

// expect User to exist
func Test_ValidateUser(t *testing.T) {
	user_id, _, err := _user_repo.ValidateUser(ValidateUserPayload{
		UserId:   _state.user_uuid,
		Password: "password",
	})

	assert.Nil(t, err)

	assert.Equal(t, _state.user_uuid, *user_id)
}

func Test_VerifyUserEmail(t *testing.T) {
	err := _user_repo.VerifyUserEmail(_state.user_uuid)
	assert.Nil(t, err)
}

func Test_GetUserByUsernameAndCheckEmailIsVerified(t *testing.T) {
	user, err := _user_repo.GetUserByUsername(_state.name)

	assert.Nil(t, err)

	assert.Equal(t, _state.user_uuid, user.ID)
	assert.Equal(t, *user.IsVerified, true)

	// would be compared to changed password
	_state.old_password_hash = user.Password
}

func Test_UpdatePassword(t *testing.T) {
	err := _user_repo.UpdateUserPassword(_state.user_uuid, "furina_doko")

	assert.Nil(t, err)
}

func Test_GetUserByIdAndCheckPassword(t *testing.T) {
	user, err := _user_repo.GetUserById(GetUserPayload{
		Id: &_state.user_uuid,
	})

	assert.Nil(t, err)

	assert.Equal(t, _state.user_uuid, user.ID)
	assert.NotEqual(t, _state.old_password_hash, user.Password)
}

func Test_GetUserByCustomId(t *testing.T) {
	user, err := _user_repo.GetUserById(GetUserPayload{
		CustomId: _state.user_uuid.String(),
	})

	assert.Nil(t, err)

	assert.Equal(t, _state.user_uuid, user.ID)
}

func Test_GetUserByProviderId(t *testing.T) {
	user, err := _user_repo.GetUserByProviderId(_state.user_uuid.String())

	assert.Nil(t, err)

	assert.Equal(t, _state.user_uuid, user.ID)
}

func Test_DeleteUser(t *testing.T) {
	err := _user_repo.DeleteUser(_state.user_uuid)
	assert.Nil(t, err)
}
