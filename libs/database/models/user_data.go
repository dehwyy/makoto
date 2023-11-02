package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider string

const (
	ProviderLocal   AuthProvider = "local"
	ProviderGoogle  AuthProvider = "google"
	ProviderGithub  AuthProvider = "github"
	ProviderDiscord AuthProvider = "discord"
)

type UserData struct {
	ID         uuid.UUID    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // @see https://www.postgresql.org/docs/current/functions-uuid.html
	Username   string       `gorm:"not null; unique; index"`
	Email      string       `gorm:"not null"`
	IsVerified *bool        `gorm:"default:false"`           // only for user.provider == local
	CustomId   string       `gorm:"not null; unique; index"` // by default, CustomId == ProviderId
	Picture    string       `gorm:"size:2048"`               // would be data64 or url
	Provider   AuthProvider `gorm:"default:local"`           // "local" | "google" | "github" | "discord"
	Role       string       `gorm:"default:user"`            // "user" | "admin"
	ProviderId string       // Id specified in OAuthProvider, if Provider==local -> same as ID
	Password   string
	Token      UserToken `gorm:"foreignKey:UserId;references:ID"`
	CreatedAt  time.Time
}
