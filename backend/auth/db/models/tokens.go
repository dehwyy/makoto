package models

type Token struct {
	ID     uint `gorm:"primaryKey"`
	Token  string
	UserId string
}
