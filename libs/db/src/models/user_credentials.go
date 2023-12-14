package models

import "github.com/google/uuid"

type UserCredentials struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // @see https://www.postgresql.org/docs/current/functions-uuid
	Usernmae string    `gorm:"unique;not null;index"`
	Email    string    `gorm:"unique;not null"`
	Password string    // hashed
}
