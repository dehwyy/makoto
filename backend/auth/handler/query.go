// @read README.md
package handler

import (
	"context"

	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *server) GetUser(ctx context.Context, input *auth.UserGetRequest) (*auth.UserDataResponse, error) {
	user, err := s.credentials_service.GetUserById(input.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &auth.UserDataResponse{
		Username: user.Username,
	}, nil
}

func (s *server) ChangePasswordByAnswer(ctx context.Context, input *auth.UserSendAnswerAndChangePasswordRequest) (*auth.UserChangePasswordResponse, error) {

	return nil, nil
}

func (s *server) ChangePassword(ctx context.Context, input *auth.UserChangePasswordRequest) (*auth.UserChangePasswordResponse, error) {
	// var username string
	// var errors *helpers.ErrorsHandler

	// go func() {
	// 	user, user_err := s.credentials_service.GetUserById(input.UserId)
	// 	if user_err != nil {
	// 		if err == nil {
	// 			err = user_err
	// 		}
	// 		return
	// 	}

	// 	username = user.Username
	// }()

	// go func() {
	// 	//! 1. Validate <old_password> to match <password> in db
	// 	// this f returns err if password doesn't match or user wasn't found
	// 	if s.credentials_service.ValidateUserPasswordById(input.UserId, input.OldPassword) != nil {
	// 		return nil, status.Errorf(codes.PermissionDenied, "wrong password")
	// 	}

	// 	//! 2. Change password
	// 	// returns arror
	// 	if s.credentials_service.UpdatePassword(input.UserId, input.NewPassword) != nil {
	// 		return nil, status.Errorf(codes.NotFound, "User was not found")
	// 	}
	// }()

	// //! 3. Update token
	// s.token_service.SignTokensAndUpdate(input.Username, input.UserId)

	return nil, nil
}
