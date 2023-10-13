package twirp

import (
	"context"
	"strings"

	"github.com/dehwyy/makoto/apps/auth/internal/oauth2"
	"github.com/dehwyy/makoto/apps/auth/internal/pipes"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
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

func (s *Server) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.AuthResponse, error) {
	s.l.Debugf("%v", req)
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
		return nil, err
	}

	// actually create
	token, err := s.token_repository.CreateOrUpdateToken(user_uuid, req.Username)

	if err != nil {
		s.l.Errorf("create token: %v", err)
		return nil, err
	}

	tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token)

	return &auth.AuthResponse{
		Username: req.Username,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthResponse, error) {
	// TODO: by Authorization header ( bearer token )

	headers, _ := tw.HTTPRequestHeaders(ctx)
	token := headers.Get("Authorization")
	// if token != "" {
	// 	return nil, nil
	// }

	// OAuth2 SignIn flow
	if oauth2_input := req.GetOauth2(); oauth2_input != nil {
		// TODO: write helper func which would return OAuth2Interface with func like `GetTokens`, `DoRequest`
		// found_user_id is userId which was found by access_token in db, would be nil if not exists
		token, found_user_id, status := s.oauth2_google.GetToken(s.parseBearerToken(oauth2_input.GetToken()), oauth2_input.GetCode())
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

		// if token == nil => no user was found => create new user + new token in db
		if found_user_id == nil {
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

		err = s.token_repository.UpdateTokenByOAuth2Token(*found_user_id, token)
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
	// either username or email should be provided
	username := credentials.GetUsername()
	email := credentials.GetEmail()
	//
	password := credentials.GetPassword()

	userId, err := s.user_repository.ValidateUser(repository.ValidateUserPayload{
		Username: username,
		Email:    email,
		Password: password,
	})

	// TODO: handle error properly
	if err != nil {
		return nil, nil
	}

	// would update it
	token, err = s.token_repository.CreateOrUpdateToken(*userId, token)
	if err != nil {
		return nil, err
	}

	tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token)

	return &auth.AuthResponse{
		Username: username,
	}, nil
}

func (s *Server) parseBearerToken(bearer_token string) (token string) {
	if bearer_token == "" {
		return
	}

	token = strings.Split(bearer_token, " ")[1]
	if len(token) < 1 {
		return
	}

	return token
}
