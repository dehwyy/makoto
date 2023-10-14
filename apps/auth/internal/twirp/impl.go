package twirp

import (
	"context"
	"errors"
	"strings"

	"github.com/dehwyy/makoto/apps/auth/internal/oauth2"
	"github.com/dehwyy/makoto/apps/auth/internal/repository"
	"github.com/dehwyy/makoto/config"
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/grpc/generated/auth"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/dehwyy/makoto/libs/middleware"
	"github.com/google/uuid"
	tw "github.com/twitchtv/twirp"
	oauth2lib "golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Server struct {
	// repository
	token_repository *repository.TokenRepository
	user_repository  *repository.UserRepository

	// oauth
	oauth2 *oauth2.OAuth2

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
		oauth2: oauth2.NewOAuth2(token_repo, config, l),

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
	var found_token *models.UserToken
	var found_user *models.UserData
	// ! `if authorization header exists -> try to auth via token
	s.l.Debugf("req %v", req)
	if req.GetEmpty() != nil {
		s.l.Debugf("empty with token value: %v", token)
		// try to find this token in db
		token_db, err := s.token_repository.GetToken(token)
		if err != nil {
			s.l.Errorf("get token: %v", err)
			return nil, tw.PermissionDenied.Error(err.Error())
		}

		user, err := s.user_repository.GetUserById(repository.GetUserPayload{
			Id: &token_db.UserId,
		})
		if err != nil {
			s.l.Errorf("get user: %v", err)
			return nil, tw.PermissionDenied.Error(err.Error())
		}

		if user.Provider == repository.ProviderLocal {
			if s.token_repository.ValidateToken(token_db.AccessToken) != nil {
				return nil, tw.Unauthenticated.Error("token expired")
			}

			// actually error should not appear heere
			if err != nil {
				return nil, tw.InternalError(err.Error())
			}

			s.l.Infof("User found by Authorization header: %v", *user)

			s.setAuthorizationHeader(ctx, token_db.AccessToken)

			return &auth.AuthResponse{
				Username: user.Username,
			}, nil
		}

		found_user = user
		found_token = token_db
	}

	// OAuth2 SignIn flow
	if oauth2_input := req.GetOauth2(); oauth2_input != nil || found_user != nil {
		s.l.Debugf("token %s", token)

		// is it ok? xd
		var oauth2_inst oauth2.OAuth2Provider
		var token_db *oauth2lib.Token
		var status oauth2.TokenStatus

		// direct request
		if oauth2_input != nil {
			oauth2_inst = s.oauth2.GetProviderInstance(oauth2_input.Provider)
			token_db, status = oauth2_inst.GetToken(token, oauth2_input.GetCode())

		} else { // only token request (it proceeds here from above (scope `if request is Empty`))
			oauth2_inst = s.oauth2.GetProviderInstance(string(found_user.Provider))
			token_db, status = oauth2_inst.GetToken(found_token.AccessToken, "")
		}

		switch status {
		case oauth2.Redirect:
			return nil, tw.NewError(tw.Unauthenticated, "provide provider credentials")
		case oauth2.InternalError:
			return nil, tw.NewError(tw.Internal, "internal error")
		}

		response, err := oauth2_inst.GetOAuth2UserByToken(token_db)
		if err != nil {
			return nil, tw.InternalErrorf("internal error %v", err.Error())
		}

		found_user, err := s.user_repository.GetUserByProviderId(response.Id)

		//  if no user was found => create new user + new token in db => return
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// generate new uuid for user
			user_uuid := uuid.New()

			// creating payload from response and other data
			createUserPayload := repository.CreateUserPayload{
				ID:       user_uuid,
				Id:       response.Id,
				Email:    response.Email,
				Username: response.Username,
				Picture:  response.Picture, // it would remove fixed size of image in this case `s96-c`
				Provider: models.AuthProvider(oauth2_inst.GetProviderName()),
				Password: "", // no password actually
			}

			err = s.user_repository.CreateUser(createUserPayload)
			if err != nil {
				s.l.Errorf("create user: %v", err)
				return nil, err
			}

			err = s.token_repository.CreateTokenByOAuth2Token(user_uuid, token_db)
			if err != nil {
				s.l.Errorf("create token: %v", err)
				return nil, err
			}

			// Set header
			tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token_db.AccessToken)
			return &auth.AuthResponse{
				Username: response.Username,
			}, err
		}

		// else if user was found => update token in db
		err = s.token_repository.UpdateTokenByOAuth2Token(found_user.ID, token_db)
		if err != nil {
			s.l.Errorf("save token: %v", err)
			return nil, err
		}

		// set header
		tw.SetHTTPResponseHeader(ctx, "Authorization", "Bearer "+token_db.AccessToken)
		return &auth.AuthResponse{
			Username: response.Username,
		}, nil
	}

	// ! By credentials
	credentials := req.GetCredentials()

	userId, err := s.user_repository.ValidateUser(repository.ValidateUserPayload{
		// ? either Username or Email would/should be provided
		Username: credentials.GetUsername(),
		Email:    credentials.GetEmail(),
		// always
		Password: credentials.GetPassword(),
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
