package pipes

import (
	"testing"
	"time"

	"github.com/dehwyy/makoto/libs/database/models"
	"golang.org/x/oauth2"
)

func Test_ToOAuth2TokenFromDbModel(t *testing.T) {
	rn := time.Now()
	tests := []struct {
		name   string
		token  *models.UserToken
		result *oauth2.Token
	}{
		{
			name: "test",
			token: &models.UserToken{
				AccessToken:  "test1",
				RefreshToken: "test2",
				Expiry:       rn,
				TokenType:    "test3",
			},
			result: &oauth2.Token{
				AccessToken:  "test1",
				RefreshToken: "test2",
				Expiry:       rn,
				TokenType:    "test3",
			},
		},
	}

	for _, tc := range tests {
		result := ToOAuth2TokenFromDbModel(tc.token)
		if result.AccessToken != tc.result.AccessToken {
			log.Fatalf("expected: %v, got: %v", tc.result.AccessToken, result.AccessToken)
		}
		if result.RefreshToken != tc.result.RefreshToken {
			log.Fatalf("expected: %v, got: %v", tc.result.RefreshToken, result.RefreshToken)
		}
		if result.Expiry != tc.result.Expiry {
			log.Fatalf("expected: %v, got: %v", tc.result.Expiry, result.Expiry)
		}
		if result.TokenType != tc.result.TokenType {
			log.Fatalf("expected: %v, got: %v", tc.result.TokenType, result.TokenType)
		}
	}

	log.Infof("Tests passed")
}
