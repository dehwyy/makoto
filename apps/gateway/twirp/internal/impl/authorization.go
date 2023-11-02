package twirp

import (
	"context"
	"fmt"

	"github.com/dehwyy/makoto/apps/gateway/services"
	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/grpc/generated/general"
	tw "github.com/twitchtv/twirp"
)

type TwirpAuthorizationService struct {
	ReadHeader             func(context.Context) (string, error)
	SetAuthorizationHeader func(context.Context, string) error

	client auth.AuthRPC
}

func NewAuthorizationService(auth_service_url string, args TwirpAuthorizationService) auth.TwirpServer {
	return auth.NewAuthRPCServer(&TwirpAuthorizationService{
		ReadHeader: args.ReadHeader,

		client: services.NewAuthorizationService(services.AuthorizationService{
			AuthorizationServiceUrl: auth_service_url,
		}),
	}, tw.WithServerPathPrefix("/authorization"))
}

func (s *TwirpAuthorizationService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	fmt.Printf("SignUp %v", req)
	response, err := s.client.SignUp(ctx, req)
	if err != nil {
		fmt.Printf("Error %v", err)
		return nil, err
	}

	fmt.Printf("Response %v", response)

	if err = s.set_token(ctx, response.Token); err != nil {
		return nil, err
	}

	new_response := &auth.AuthResponse{
		UserId:   response.UserId,
		Username: response.Username,
	}
	return new_response, nil
}

func (s *TwirpAuthorizationService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	response, err := s.client.SignIn(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpAuthorizationService) IsUniqueUsername(ctx context.Context, req *auth.IsUniqueUsernamePayload) (*auth.IsUnique, error) {
	response, err := s.client.IsUniqueUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpAuthorizationService) VerifyUserEmail(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	response, err := s.client.VerifyUserEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpAuthorizationService) ChangePassword(ctx context.Context, req *auth.ChangePasswordPayload) (*general.IsSuccess, error) {
	response, err := s.client.ChangePassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpAuthorizationService) Logout(ctx context.Context, req *general.UserId) (*general.IsSuccess, error) {
	response, err := s.client.Logout(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TwirpAuthorizationService) set_token(ctx context.Context, token string) error {
	return tw.SetHTTPResponseHeader(ctx, "Authorization", fmt.Sprintf("Bearer %s", token))
}
