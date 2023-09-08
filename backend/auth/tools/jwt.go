package tools

import (
	"fmt"
	"time"

	"github.com/dehwyy/Makoto/backend/auth/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtKey = config.GetOptionByKey("jwt.key")
)

type Jwt struct{}

type JwtPayload struct {
	Username string
	UserId   string
}

func NewJwt() *Jwt {
	return new(Jwt)
}

func newJwtToken(payload JwtPayload, exp_minutes int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload.Username,
		"userId":   payload.UserId,
		"exp":      time.Now().UTC().Add(time.Duration(exp_minutes) * time.Minute).Unix(),
	})
	signed_token, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return signed_token, nil
}

func (j *Jwt) NewRefreshToken(payload JwtPayload) (string, error) {
	// exp_minutes is equals to 14 days 60(minutes) * 24(hours) * 14(days)
	return newJwtToken(payload, 60*24*14)
}

func (j *Jwt) NewAccessToken(payload JwtPayload) (string, error) {
	return newJwtToken(payload, 30)
}

func (j *Jwt) ValidateJwtToken(token_string string) error {
	// parse token
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// validate algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	// clarify whether parse is succeed
	if err != nil {
		return err
	}

	// Expect claims to be
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// getting expiration time and computing
	// whether token's exp date is after the now date
	// if it  already expired => return error
	exp := int64(claims["exp"].(float64))
	isExpired := time.Now().After(time.Unix(exp, 0))
	if isExpired {
		return fmt.Errorf("token is expired")
	}

	return nil
}
