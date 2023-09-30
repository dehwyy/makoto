package models

type Tag struct {
	Id    uint `gorm:"primaryKey;->"`
	Name  string
	Users []*Word `gorm:"many2many:word_tags;"`
}
