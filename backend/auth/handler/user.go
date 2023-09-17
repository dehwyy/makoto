package handler

import (
	"context"

	auth "github.com/dehwyy/Makoto/backend/grpc/gen/auth/go/proto"
)

func (s *server) SignUp(ctx context.Context, in *auth.UserSignUpRequest) (*auth.UserResponse, error) {
	// Creating user payload
	user_payload := s.credentials_service.CreateUserPayload(in.Username, in.Password, in.Question, in.Answer)

	// Creating userj
	user_id, err := s.credentials_service.CreateUser(user_payload)
	if err != nil {
		s.log.Errorf("Error creating user: %v", err)
		return nil, err
	}

	// Generating tokens based on user_id and username
	access_token, refresh_token := s.token_service.SignTokensAndCreate(in.Username, user_id)

	return &auth.UserResponse{
		UserId:       user_id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil

}

func (s *server) SignIn(ctx context.Context, in *auth.UserSignInRequest) (*auth.UserResponse, error) {
	userId, err := s.credentials_service.ValidateUser(in.Username, in.Password)

	if err != nil {
		s.log.Errorf("Error validating user: %v", err)
		return nil, err
	}

	access_token, refresh_token := s.token_service.SignTokensAndUpdate(in.Username, userId)

	return &auth.UserResponse{
		UserId:       userId,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *server) ValidateAuth(ctx context.Context, in *auth.AccessToken) (*auth.ValidateAuthResponse, error) {
	// Validate access_token token sign and encrypt method
	userId, username, isValid := s.token_service.ValidateToken(in.AccessToken)

	if !isValid {
		return &auth.ValidateAuthResponse{
			IsOk: false,
		}, nil
	}

	acccess_token, refresh_token := s.token_service.SignTokensAndUpdate(username, userId)

	return &auth.ValidateAuthResponse{
		IsOk:         true,
		UserId:       userId,
		RefreshToken: refresh_token,
		AccessToken:  acccess_token,
	}, nil
}
