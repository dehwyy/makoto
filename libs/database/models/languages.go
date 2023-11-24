package models

type Language struct {
	ID    uint        `gorm:"primarykey"`
	Lang  string      `gorm:"not null; unique"`
	Users []*UserInfo `gorm:"many2many:users_languages;"`
}
