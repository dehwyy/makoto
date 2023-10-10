package models

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	ID           uint      `gorm:"primarykey"`
	AccessToken  string    `gorm:"not null;index"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `gorm:"not null"`
	// @see https://pkg.go.dev/golang.org/x/oauth2#Token
	UserId uuid.UUID
}
