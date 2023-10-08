package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model

	// @see https://pkg.go.dev/golang.org/x/oauth2#Token
	AccessToken  string    `gorm:"not null;index"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `json:"expiry,omitempty"`
}
