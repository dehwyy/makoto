// @read README.md
package handler

import (
	"context"
	"sync"

	"github.com/dehwyy/Makoto/backend/auth/dto"
	"github.com/dehwyy/Makoto/backend/auth/helpers"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *server) GetQuestion(ctx context.Context, input *auth.UserGetQuestionRequest) (*auth.UserQuestionResponse, error) {
	question, err := s.credentials_service.GetQuestion(input.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &auth.UserQuestionResponse{
		Question: question,
	}, nil
}

func (s *server) GetUserById(ctx context.Context, input *auth.UserGetRequest) (*auth.UserDataResponse, error) {
	user, err := s.credentials_service.GetUserById(input.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &auth.UserDataResponse{
		Username: user.Username,
	}, nil
}

func (s *server) changePasswordTemplate(ctx context.Context, input *dto.UpdatePassword) (*auth.UserChangePasswordResponse, error) {
	//! 0. Find usermae by provided ID
	var username string
	err := helpers.NewErrorsHandler()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		user, e := s.credentials_service.GetUserById(input.UserId)

		wg.Done()
		if e != nil {
			err.SetErrorIfNotAlready(e)
			return
		}

		username = user.Username
	}()

	//! 1.Validate <old_password> to match <password> in db
	go func() {
		// this f returns err if password doesn't match or user wasn't found
		if input.ValidateFunc() != nil {
			e := status.Errorf(codes.PermissionDenied, "wrong password")
			err.SetErrorIfNotAlready(e)
		}

		wg.Done()
	}()

	wg.Wait()

	if e := err.GetError(); e != nil {
		return nil, e
	}

	// *declare tokens
	var access, refresh string

	// Create transaction for ChangePassword & UpdateToken actions
	e := s.db.Transaction(func(tx *gorm.DB) error {
		//! 2.Change password
		if s.credentials_service.UpdatePassword(input.UserId, input.NewPassword) != nil {
			return status.Errorf(codes.NotFound, "User was not found")
		}

		//! 3.Update token
		access, refresh = s.token_service.SignTokensAndUpdate(username, input.UserId)

		return nil
	})

	if e != nil {
		return nil, e
	}

	return &auth.UserChangePasswordResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *server) ChangePasswordByAnswer(ctx context.Context, input *auth.UserSendAnswerAndChangePasswordRequest) (*auth.UserChangePasswordResponse, error) {
	return s.changePasswordTemplate(ctx, &dto.UpdatePassword{
		UserId:      input.UserId,
		NewPassword: input.NewPassword,
		ValidateFunc: func() error {
			return s.credentials_service.ValidateUserAnswer(input.UserId, input.Answer)
		},
	})
}

func (s *server) ChangePassword(ctx context.Context, input *auth.UserChangePasswordRequest) (*auth.UserChangePasswordResponse, error) {
	return s.changePasswordTemplate(ctx, &dto.UpdatePassword{
		UserId:      input.UserId,
		NewPassword: input.NewPassword,
		ValidateFunc: func() error {
			return s.credentials_service.ValidateUserPassword(input.UserId, input.OldPassword)
		},
	})
}
