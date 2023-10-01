package models

type Tag struct {
	Id   uint32 `gorm:"primaryKey;"`
	Text string
}
