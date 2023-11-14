package models

import (
	"time"

	"github.com/google/uuid"
)

type UserInfo struct {
	ID              uint `gorm:"primarykey"`
	UserId          uuid.UUID
	Picture         string // would be data64 or url
	Description     string
	BackgroundDark  string      // color in hex format - for dark mode
	BackgroundLight string      // for light mode
	Languages       []*Language `gorm:"many2many:users_languages"`
	CreatedAt       time.Time
}
