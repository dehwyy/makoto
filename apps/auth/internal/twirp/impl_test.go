package twirp

import (
	"context"
	"errors"
	default_logger "log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dehwyy/makoto/libs/config"
	"github.com/dehwyy/makoto/libs/database"
	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	tw "github.com/twitchtv/twirp"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type _test_state struct {
	db     *gorm.DB
	userId string
}

var (
	_log   = logger.New()
	_state = &_test_state{}
)

func Test_init(t *testing.T) {
	db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", _log)
	db = db.Session(&gorm.Session{
		Logger: gorm_logger.New(default_logger.New(os.Stdout, "\r\n", default_logger.LstdFlags), gorm_logger.Config{
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})

	_state.db = db

	db.Exec("DELETE FROM user_tokens")
	db.Exec("DELETE FROM user_data")
}

func _test_new_server() map[string]auth.AuthRPC {
	handler := NewTwirpServer(_state.db, &config.Config{
		JwtSecret: "secret",
	}, _log)

	server := httptest.NewServer(handler)
	clients := map[string]auth.AuthRPC{
		"proto": auth.NewAuthRPCProtobufClient(server.URL, http.DefaultClient),
		"json":  auth.NewAuthRPCJSONClient(server.URL, http.DefaultClient),
	}

	return clients
}

func _chech_twirp_error(t *testing.T, err error, expect tw.ErrorCode) tw.Error {
	// @see https://twitchtv.github.io/twirp/docs/errors
	var twerr tw.Error
	errors.As(err, &twerr)
	assert.Equal(t, twerr.Code(), expect)

	return twerr
}

func Test_SignUp(t *testing.T) {
	var err error
	var twerr tw.Error

	client := _test_new_server()["proto"]

	// 1. success case
	res, err := client.SignUp(context.Background(), &auth.SignUpRequest{
		Username: "dehwyy_tester",
		Email:    "dehwyy_tester@example.com",
		Password: "password",
	})

	assert.Nil(t, err)

	assert.NotEmpty(t, res.Token)
	assert.NotEmpty(t, res.UserId)
	assert.Equal(t, res.Username, "dehwyy_tester")
	assert.Equal(t, res.IsCreated, true)

	_state.userId = res.UserId

	// 2. UserAlreadyExistsError
	res, err = client.SignUp(context.Background(), &auth.SignUpRequest{
		Username: "dehwyy_tester",
		Email:    "dehwyy_tester@example.com",
		Password: "password",
	})
	assert.Nil(t, res)
	_chech_twirp_error(t, err, tw.InvalidArgument)

	// 3. With validation fields (non-empty fields)
	res, err = client.SignUp(context.Background(), &auth.SignUpRequest{
		Username: "",
		Email:    "mail",
		Password: "password",
	})

	assert.Nil(t, res)
	twerr = _chech_twirp_error(t, err, tw.InvalidArgument)
	assert.Contains(t, twerr.Msg(), "username")

	res, err = client.SignUp(context.Background(), &auth.SignUpRequest{
		Username: "dehwyy_tester",
		Email:    "",
		Password: "password",
	})

	assert.Nil(t, res)
	twerr = _chech_twirp_error(t, err, tw.InvalidArgument)
	assert.Contains(t, twerr.Msg(), "email")

	res, err = client.SignUp(context.Background(), &auth.SignUpRequest{
		Username: "dehwyy_tester",
		Email:    "dehwyy_tester@example.com",
		Password: "",
	})

	assert.Nil(t, res)
	twerr = _chech_twirp_error(t, err, tw.InvalidArgument)
	assert.Contains(t, twerr.Msg(), "password")
}

func Test_SignIn(t *testing.T) {
	client := _test_new_server()["proto"]

	// success
	// 1. via password + username
	res, err := client.SignIn(context.Background(), &auth.SignInRequest{
		AuthMethod: &auth.SignInRequest_Credentials{
			Credentials: &auth.SignInRequest_UserCredentials{
				Password: "password",
				UniqueIdentifier: &auth.SignInRequest_UserCredentials_Username{
					Username: "dehwyy_tester",
				},
			},
		},
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, res.Token)
	assert.Equal(t, res.Username, "dehwyy_tester")
	assert.Equal(t, res.UserId, _state.userId)
	assert.Equal(t, res.IsCreated, false)

	// 2. via password + email
	res, err = client.SignIn(context.Background(), &auth.SignInRequest{
		AuthMethod: &auth.SignInRequest_Credentials{
			Credentials: &auth.SignInRequest_UserCredentials{
				Password: "password",
				UniqueIdentifier: &auth.SignInRequest_UserCredentials_Email{
					Email: "dehwyy_tester@example.com",
				},
			},
		},
	})

	_log.Infof("res: %v", res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res.Token)
	assert.Equal(t, res.Username, "dehwyy_tester")
	assert.Equal(t, res.UserId, _state.userId)
	assert.Equal(t, res.IsCreated, false)

	login_token := res.Token

	// 3. via token
	res, err = client.SignIn(context.Background(), &auth.SignInRequest{
		AuthMethod: &auth.SignInRequest_Token{
			Token: login_token,
		},
	})

	_log.Infof("res: %v", res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res.Token)
	assert.Equal(t, res.Username, "dehwyy_tester")
	assert.Equal(t, res.UserId, _state.userId)
	assert.Equal(t, res.IsCreated, false)

	// fail
	// 1. Token Errors
	// Invalid token
	res, err = client.SignIn(context.Background(), &auth.SignInRequest{
		AuthMethod: &auth.SignInRequest_Token{
			Token: "invalid_token",
		},
	})

	assert.Nil(t, res)
	twerr := _chech_twirp_error(t, err, tw.PermissionDenied)
	assert.Contains(t, twerr.Msg(), "not found")
}

func Test_IsUniqueUsername(t *testing.T) {
	client := _test_new_server()["proto"]
	res, err := client.IsUniqueUsername(context.Background(), &auth.IsUniqueUsernamePayload{
		Username: "dehwyy_tester",
	})

	assert.Nil(t, err)
	assert.Equal(t, res.IsUnique, false)

	res, err = client.IsUniqueUsername(context.Background(), &auth.IsUniqueUsernamePayload{
		Username: "dehwyy_tester_doko",
	})

	assert.Nil(t, err)
	assert.Equal(t, res.IsUnique, true)

	//
}

func Test_ChangePassword(t *testing.T) {
	client := _test_new_server()["proto"]

	// success
	res, err := client.ChangePassword(context.Background(), &auth.ChangePasswordPayload{
		UserId:      _state.userId,
		OldPassword: "password",
		NewPassword: "new_password",
	})

	assert.Nil(t, err)
	assert.Equal(t, res.IsSuccess, true)

	// fail
	// 1. Wrong password
	res, err = client.ChangePassword(context.Background(), &auth.ChangePasswordPayload{
		UserId:      _state.userId,
		OldPassword: "wrong_password",
		NewPassword: "new_password",
	})

	assert.NotNil(t, err)
	twerr := _chech_twirp_error(t, err, tw.Unauthenticated)
	assert.Contains(t, twerr.Msg(), "password")
	assert.Nil(t, res)

	// 2. UserNotFound
	res, err = client.ChangePassword(context.Background(), &auth.ChangePasswordPayload{
		UserId:      uuid.New().String(),
		OldPassword: "password",
		NewPassword: "new_password",
	})

	assert.NotNil(t, err)
	twerr = _chech_twirp_error(t, err, tw.NotFound)
	assert.Contains(t, twerr.Msg(), "user")
	assert.Nil(t, res)

	// 3. InvalidUserId (not uuid)
	res, err = client.ChangePassword(context.Background(), &auth.ChangePasswordPayload{
		UserId:      "not_uuid",
		OldPassword: "password",
		NewPassword: "new_password",
	})

	assert.NotNil(t, err)
	twerr = _chech_twirp_error(t, err, tw.InvalidArgument)
	assert.Contains(t, twerr.Msg(), "userId")
	assert.Nil(t, res)
}

func Test_Logout(t *testing.T) {
	client := _test_new_server()["proto"]

	// success
	res, err := client.Logout(context.Background(), &general.UserId{
		UserId: _state.userId,
	})

	assert.Nil(t, err)
	assert.Equal(t, res.IsSuccess, true)

	// fail
	// 1. UserNotFound
	res, err = client.Logout(context.Background(), &general.UserId{
		UserId: uuid.New().String(),
	})

	assert.NotNil(t, err)
	twerr := _chech_twirp_error(t, err, tw.NotFound)
	assert.Contains(t, twerr.Msg(), "token")
	assert.Nil(t, res)

	// 2. InvalidUserId (not uuid)
	res, err = client.Logout(context.Background(), &general.UserId{
		UserId: "not_uuid",
	})

	assert.NotNil(t, err)
	twerr = _chech_twirp_error(t, err, tw.InvalidArgument)
	assert.Contains(t, twerr.Msg(), "userId")
	assert.Nil(t, res)
}
