package models

type Tag struct {
	Id    uint `gorm:"primaryKey;->"`
	Text  string
	Users []*Word `gorm:"many2many:word_tags;"`
}
