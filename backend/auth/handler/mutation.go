// @read README.md
package handler

import (
	"context"

	"github.com/dehwyy/Makoto/backend/auth/dto"
	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) SignUp(ctx context.Context, in *auth.UserSignUpRequest) (*auth.UserResponse, error) {
	// Creating user payload
	user_payload := dto.CreateUser{
		Username: in.Username,
		Password: in.Password,
		Question: in.Question,
		Answer:   in.Answer,
	}

	// Creating userj
	user_id, err := s.credentials_service.CreateUser(user_payload)
	if err != nil {
		s.log.Errorf("Error creating user: %v", err)
		return nil, status.Errorf(codes.Internal, "error creating user")
	}

	// Generating tokens based on user_id and username
	access_token, refresh_token := s.token_service.SignTokensAndCreate(in.Username, user_id)

	return &auth.UserResponse{
		UserId:       user_id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *server) SignOut(ctx context.Context, in *auth.UserSignOutRequest) (*auth.Nil, error) {
	err := s.credentials_service.RemoveToken(in.UserId)
	if err != nil {
		return nil, err
	}

	return &auth.Nil{}, nil
}

func (s *server) SignIn(ctx context.Context, in *auth.UserSignInRequest) (*auth.UserResponse, error) {
	userId, err := s.credentials_service.ValidateUser(in.Username, in.Password)

	if err != nil {
		s.log.Errorf("Error validating user: %v", err)
		if userId == "403" {
			return nil, status.Errorf(codes.PermissionDenied, "invalid password")
		}
		return nil, status.Errorf(codes.NotFound, "user with such username was not found")
	}

	access_token, refresh_token := s.token_service.SignTokensAndUpdate(in.Username, userId)

	return &auth.UserResponse{
		UserId:       userId,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *server) ValidateAuth(ctx context.Context, in *auth.AccessToken) (*auth.UserResponse, error) {
	// Validate access_token token sign and encrypt method
	userId, username, isValid := s.token_service.ValidateToken(in.AccessToken)

	if !isValid {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	acccess_token, refresh_token := s.token_service.SignTokensAndUpdate(username, userId)

	return &auth.UserResponse{
		UserId:       userId,
		RefreshToken: refresh_token,
		AccessToken:  acccess_token,
	}, nil
}
