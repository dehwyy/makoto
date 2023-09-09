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
	}

	// Generating tokens based on user_id and username
	access_token, refresh_token := s.token_service.SignTokensAndSave(in.Username, user_id)

	return &auth.UserResponse{
		UserId:       user_id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil

}
