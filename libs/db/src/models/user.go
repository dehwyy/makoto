package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserCredentials struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // @see https://www.postgresql.org/docs/current/functions-uuid
	Username  string    `gorm:"unique;not null;index"`
	Email     string    `gorm:"unique;not null"`
	Password  string    // hashed password
	CreatedAt time.Time

	// relations
	Tokens UserTokens `gorm:"foreignKey:UserId;references:ID"`
	OAuth  UserOauth  `gorm:"foreignKey:UserId;references:ID"`
}

type UserTokens struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId       uuid.UUID      `gorm:"not null;index"` // foreign key
	AccessToken  pq.StringArray `gorm:"type:text[]"`
	RefreshToken string         `gorm:"not null"`
	Expiry       time.Time      `gorm:"not null"`
}

type UserOauth struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId    uuid.UUID `gorm:"not null;index"` // foreign key18
	GoogleId  string    `gorm:"unique"`
	GithubId  string    `gorm:"unique"`
	DiscordId string    `gorm:"unique"`
}
