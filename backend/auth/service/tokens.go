package service

import "github.com/dehwyy/Makoto/backend/auth/logger"

type jwtPayload struct {
	Username string
	UserId   string
}

type JwtHandler interface {
	// (payload jwtPayload) => (token: string, error: error)
	NewRefreshToken(jwtPayload) (string, error)
	// same as previous
	NewAccessToken(jwtPayload) (string, error)
	// (token: string) => error (if valid => nil)
	ValidateJwtToken(string) error
}

type token_service struct {
	jwt JwtHandler
	l   logger.AppLogger
}

func NewTokenService(jwt JwtHandler) *token_service {
	return &token_service{
		jwt: jwt,
	}
}

func (t *token_service) newJwtPayload(username, userId string) jwtPayload {
	return jwtPayload{
		Username: username,
		UserId:   userId,
	}
}

// Signing access and refresh tokens; The refresh one would be saved to db; Returns (ACCESS_TOKEN, REFRESH_TOKEN).
func (t *token_service) SignTokensAndSave(username, userId string) (string, string) {
	payload := t.newJwtPayload(username, userId)

	refresh_token, err := t.jwt.NewRefreshToken(payload)
	if err != nil {
		t.l.Errorf("Error creating refresh token: %v", err)
	}
	// TODO: save to db
	t.l.Infof("Generated refresh token: %v", refresh_token)

	access_token, err := t.jwt.NewAccessToken(payload)
	if err != nil {
		t.l.Errorf("Error creating access token: %v", err)
	}
	t.l.Infof("Generated access token: %v", access_token)

	return access_token, refresh_token
}
