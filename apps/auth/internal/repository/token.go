package repository

import (
	"errors"

	"github.com/dehwyy/makoto/apps/auth/internal/utils"
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type TokenRepositoryReadonly interface {
	GetToken(access_token string) (*models.UserToken, error)
}

type TokenRepository struct {
	db  *gorm.DB
	jwt *utils.Jwt
	l   logger.Logger
}

// NewTokenRepository initializes a new instance of TokenRepository.
//
// It takes in a *gorm.DB as the first parameter which represents the database connection,
// a logger.Logger as the second parameter which is used for logging,
// and a string jwt_secret as the third parameter which is the secret key for JWT token generation.
//
// It returns a *TokenRepository.
func NewTokenRepository(db *gorm.DB, l logger.Logger, jwt_secret string) *TokenRepository {
	return &TokenRepository{
		db:  db,
		jwt: utils.NewJwt(jwt_secret),
		l:   l,
	}
}

// GetToken retrieves a user token from the TokenRepository based on the provided access token.
//
// Parameters:
// - access_token: The access token used to search for the user token.
//
// Returns:
// - *models.UserToken: The user token found based on the access token.
// - error: An error if there was a problem retrieving the user token.
func (t *TokenRepository) GetToken(access_token string) (*models.UserToken, error) {
	var token models.UserToken

	res := t.db.Model(&models.UserToken{}).Where("access_token = ?", access_token).First(&token)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return &token, res.Error
}

// CreateToken creates an token for the given user ID and username.
//
// Parameters:
// - userId: the UUID of the user.
// - username: the username of the user.
//
// Returns:
// - token: the access token.
// - err: an error if the token creation fails.
func (t *TokenRepository) CreateToken(userId uuid.UUID, username string) (token string, err error) {
	jwtPayload := utils.JwtPayload{
		Username: username,
		UserId:   userId.String(),
	}

	token, exp, err_access := t.jwt.NewAccessToken(jwtPayload)
	refresh_token, _, err_refresh := t.jwt.NewRefreshToken(jwtPayload)
	if err_access != nil || err_refresh != nil {
		t.l.Errorf("Failder to create token %v, %v", err_access, err_refresh)
		return "", errors.New("failed to create token")
	}

	return token, t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token,
		RefreshToken: refresh_token,
		Expiry:       exp,
		TokenType:    "Bearer",
	}).Error
}

// UpdateToken updates the token for a given user ID.
//
// Parameters:
// - userId: the ID of the user to update the token for.
//
// Returns:
// - token: the updated token string.
// - err: an error if the token update fails.
func (t *TokenRepository) UpdateToken(userId uuid.UUID) (token string, err error) {
	token, exp, err := t.jwt.NewAccessToken(utils.JwtPayload{
		UserId: userId.String(),
	})

	if err != nil {
		return "", err
	}

	return token, t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken: token,
		Expiry:      exp,
	}).Error
}

// ValidateToken validates a token.
//
// It takes a token string as a parameter.
// It returns an error.
func (t *TokenRepository) ValidateToken(token string) error {
	_, err := t.jwt.ValidateJwtToken(token)
	return err
}

// DeleteToken removes a token from the TokenRepository.
//
// It takes a token string as a parameter.
// It returns an error.
func (t *TokenRepository) DeleteToken(userId uuid.UUID) error {
	return t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Delete(&models.UserToken{}).Error
}

// ! if authorization method is OAuth2
func (t *TokenRepository) UpdateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Updates(&models.UserToken{
		AccessToken: token.AccessToken,
		Expiry:      token.Expiry,
	}).Error
}

func (t *TokenRepository) CreateTokenByOAuth2Token(userId uuid.UUID, token *oauth2.Token) error {
	return t.db.Model(&models.UserToken{}).Create(&models.UserToken{
		UserId:       userId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
	}).Error
}

func (t *TokenRepository) DeleteTokenByUserId(userId uuid.UUID) error {
	res := t.db.Model(&models.UserToken{}).Where("user_id = ?", userId).Delete(&models.UserToken{})

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return res.Error
}
