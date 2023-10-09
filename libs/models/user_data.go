package models

import (
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // @see https://www.postgresql.org/docs/current/functions-uuid.html
	Username  string    `gorm:"not null; unique; index"`
	Email     string    `gorm:"not null; unique"`
	CustomId  string    `gorm:"not null; unique"`
	Picture   string    `gorm:"size:2048"`     // would be data64
	Provider  string    `gorm:"default:local"` // "local" | "Google"
	Role      string    `gorm:"default:user"`  // "user" | "admin"
	Password  string
	TokenID   uint // @see https://gorm.io/docs/belongs_to.html
	Token     UserToken
	CreatedAt time.Time
}
