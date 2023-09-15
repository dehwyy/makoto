package models

import "time"

type Credentials struct {
	ID           uint   `gorm:"primaryKey;->"`
	UniqueUserId string `gorm:"unique;index"`
	Username     string `gorm:"unique;index"`
	Password     string
	Question     string
	Answer       string
	CreatedAt    time.Time
	token        Token
}
