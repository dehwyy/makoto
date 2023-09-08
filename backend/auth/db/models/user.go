package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;->"`
	UserId    string `gorm:"unique;index"`
	Username  string `gorm:"unique;index"`
	Password  string
	Question  string
	Answer    string
	CreatedAt time.Time
}
