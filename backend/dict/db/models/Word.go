package models

type Word struct {
	Id     uint32 `gorm:"primaryKey;"`
	UserId string `gorm:"index"`
	Word   string
	Value  string
	Extra  string
	Tags   []Tag `gorm:"many2many:word_tags;"`
}
