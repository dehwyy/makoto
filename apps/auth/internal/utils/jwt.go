package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	jwt_key string
}

type JwtPayload struct {
	Username string
	UserId   string
}

func NewJwt(jwt_key string) *Jwt {
	return &Jwt{
		jwt_key: jwt_key,
	}
}

func (j *Jwt) NewRefreshToken(payload JwtPayload) (token string, exp time.Time, err error) {
	// exp_minutes is equals to 30 days 60(minutes) * 24(hours) * 30(days)
	return j.newJwtToken(payload, 60*24*30)
}

func (j *Jwt) NewAccessToken(payload JwtPayload) (token string, exp time.Time, err error) {
	return j.newJwtToken(payload, 30)
}

func (j *Jwt) ValidateJwtToken(token_string string) (*JwtPayload, error) {
	// parse token
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// validate algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.jwt_key), nil
	})
	// clarify whether parse is succeed
	if err != nil {
		return nil, err
	}

	// Expect claims to be
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// getting expiration time and computing
	// whether token's exp date is after the now date
	// if it  already expired => return error
	exp := int64(claims["exp"].(float64))
	isExpired := time.Now().After(time.Unix(exp, 0))
	if isExpired {
		return nil, fmt.Errorf("token is expired")
	}

	return &JwtPayload{
		Username: claims["username"].(string),
		UserId:   claims["userId"].(string),
	}, nil
}

func (j *Jwt) newJwtToken(payload JwtPayload, exp_minutes int) (signed_token string, exp time.Time, err error) {
	exp = time.Now().UTC().Add(time.Duration(exp_minutes) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload.Username,
		"userId":   payload.UserId,
		"exp":      exp.Unix(),
	})
	signed_token, err = token.SignedString([]byte(j.jwt_key))
	if err != nil {
		return "", exp, err
	}

	return signed_token, exp, nil
}
