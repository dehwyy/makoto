package twirp

import (
	"context"
	"errors"
	"strings"

	"github.com/dehwyy/makoto/apps/auth/internal/oauth2"
	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/middleware"
	"github.com/google/uuid"
	tw "github.com/twitchtv/twirp"
	"gorm.io/gorm"
)

type Server struct {
	// repository
	token_repository *repository.TokenRepository
	user_repository  *repository.UserRepository

	// oauth
	oauth2_google *oauth2.Google

	//
	l logger.Logger
}

func NewTwirpServer(db *gorm.DB, config config.Config, l logger.Logger) auth.TwirpServer {
	token_repo := repository.NewTokenRepository(db, l, config.JwtSecret)
	user_repo := repository.NewUserRepository(db, l)

	return auth.NewAuthServer(&Server{
		// repo
		token_repository: token_repo,
		user_repository:  user_repo,

		// oauth
		oauth2_google: oauth2.NewGoogleOAuth2(config.Oauth2.Google.Id, config.Oauth2.Google.Secret, config.Oauth2.Google.RedirectURL, token_repo, l),

		// logger
		l: l,
	})
}

// Only for credentials
func (s *Server) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	user_uuid := uuid.New()

	if err := s.user_repository.CreateUser(repository.CreateUserPayload{
		ID:       user_uuid,
		Id:       user_uuid.String(),
		Email:    req.Email,
		Username: req.Username,
		Picture:  "",
		Password: req.Password,
		Provider: repository.ProviderLocal,
	}); err != nil {
		s.l.Errorf("create user: %v", err)
		return nil, tw.InternalError(err.Error())
	}

	token, err := s.token_repository.CreateToken(user_uuid, req.Username)
	if err != nil {
		s.l.Errorf("create token: %v", err)
		return nil, tw.InternalError(err.Error())
	}

	s.setAuthorizationHeader(ctx, token)

	return &auth.AuthResponse{
		Username: req.Username,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {

	token := s.parseBearerToken(middleware.WithAuthorizationMiddlewareRead(ctx))

	// ! `if authorization header exists -> try to auth via token
	s.l.Debugf("req %v", req)
	if req.GetEmpty() != nil {
		s.l.Debugf("HERE")
		// try to find this token in db
		found_token, err := s.token_repository.GetToken(token)
		if err != nil {
			s.l.Errorf("get token: %v", err)
			return nil, tw.PermissionDenied.Error(err.Error())
		}

		user, err := s.user_repository.GetUserById(repository.GetUserPayload{
			Id: &found_token.UserId,
		})

		if user.Provider == repository.ProviderLocal {
			if s.token_repository.ValidateToken(found_token.AccessToken) != nil {
				return nil, tw.Unauthenticated.Error("token expired")
			}

			// actually error should not appear heere
			if err != nil {
				return nil, tw.InternalError(err.Error())
			}

			s.l.Infof("User found by Authorization header: %v", *user)

			s.setAuthorizationHeader(ctx, found_token.AccessToken)

			return &auth.AuthResponse{
				Username: user.Username,
			}, nil
		}

		token = found_token.AccessToken
	}

	// OAuth2 SignIn flow
	if oauth2_input := req.GetOauth2(); oauth2_input != nil || token != "" {
		// TODO: write helper func which would return OAuth2Interface with func like `GetTokens`, `DoRequest`
		// found_user_id is userId which was found by access_token in db, would be nil if not exists
		s.l.Debugf("token %s", token)
		token, status := s.oauth2_google.GetToken(token, oauth2_input.GetCode())

		switch status {
		case oauth2.Redirect:
			return nil, tw.NewError(tw.Unauthenticated, "provide google credentials")
		case oauth2.InternalError:
			return nil, tw.NewError(tw.Internal, "internal error")
		}

		res, err := s.oauth2_google.DoRequest(oauth2.GoogleProfile, token)
		if err != nil {
			s.l.Errorf("request: %v", err)
			return nil, err
		}

		var GoogleResponse struct {
			Id      string `json:"id"`
			Email   string `json:"email"`
			Name    string `json:"name"`
			Picture string `json:"picture"`
			// email:dehwyy@gmail.com
			// given_name:dehwyy
			// id: 103623406957472659690
			// name:dehwyy
			// picture: `https://lh3.googleusercontent.com/a/ACg8ocLE4oqn1c6KC1jgzJB3vL3hhJBDEKxINbHfQmG34Ubrozk=s96-c`
		}

		if err := pipes.Body2Struct(res.Body, &GoogleResponse); err != nil {
			s.l.Errorf("pipes res.body %v", err)
			return nil, err
		}

		found_user, err := s.user_repository.GetUserByProviderId(GoogleResponse.Id)

		//  if no user was found => create new user + new token in db => return
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// generate new uuid for user
			user_uuid := uuid.New()

			// creating payload from response and other data
			createUserPayload := repository.CreateUserPayload{
				ID:       user_uuid,
				Id:       GoogleResponse.Id,
				Email:    GoogleResponse.Email,
				Username: GoogleResponse.Name,
				Picture:  strings.Split(GoogleResponse.Picture, "=")[0], // it would remove fixed size of image in this case `s96-c`
				Provider: repository.ProviderGoogle,
				Password: "", // no password actually
			}

			err = s.user_repository.CreateUser(createUserPayload)
			if err != nil {
				s.l.Errorf("create user: %v", err)
				return nil, err
			}

			err = s.token_repository.CreateTokenByOAuth2Token(user_uuid, token)
			if err != nil {
				s.l.Errorf("create token: %v", err)
				return nil, err
			}

			// Set header
			tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token.AccessToken)
			return &auth.AuthResponse{
				Username: GoogleResponse.Name,
			}, err
		}

		// else if user was found => update token in db
		err = s.token_repository.UpdateTokenByOAuth2Token(found_user.ID, token)
		if err != nil {
			s.l.Errorf("save token: %v", err)
			return nil, err
		}

		// set header
		tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token.AccessToken)
		return &auth.AuthResponse{
			Username: GoogleResponse.Name,
		}, nil
	}

	// ! By credentials
	credentials := req.GetCredentials()

	userId, err := s.user_repository.ValidateUser(repository.ValidateUserPayload{
		// ? either Username or Email would/should be provided
		Username: credentials.GetUsername(),
		Email:    credentials.GetEmail(),
		// always
		Password: credentials.Password,
	})

	if errors.Is(err, repository.USER_NOT_FOUND) {
		return nil, tw.NewError(tw.Unauthenticated, "user with provided credentials wasn't found")

	} else if errors.Is(err, repository.USER_WRONG_PASSWORD) {
		return nil, tw.NewError(tw.Unauthenticated, "wrong password")

	} // would not return unnamed error => no simple check for nil

	token, err = s.token_repository.UpdateToken(*userId, token)
	if err != nil {
		return nil, tw.InternalError(err.Error())
	}

	// settings http authorization header
	s.setAuthorizationHeader(ctx, token)

	return &auth.AuthResponse{
		Username: credentials.GetUsername(),
	}, nil
}

func (s *Server) parseBearerToken(bearer_token string) (token string) {
	split_token := strings.Split(bearer_token, " ")
	if len(split_token) < 2 {
		return
	}

	token = strings.Split(bearer_token, " ")[1]
	if len(token) < 1 {
		return
	}

	return token
}

func (s *Server) setAuthorizationHeader(ctx context.Context, token string) {
	tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token)
}
