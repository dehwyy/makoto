package pipes

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"golang.org/x/oauth2"
)

func ToOAuth2TokenFromDbModel(dbToken *models.UserToken) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  dbToken.AccessToken,
		RefreshToken: dbToken.RefreshToken,
		Expiry:       dbToken.Expiry,
		TokenType:    dbToken.TokenType,
	}
}
