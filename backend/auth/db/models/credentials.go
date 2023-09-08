package models

type Credentials struct {
	ID     uint `gorm:"primaryKey"`
	Token  string
	UserId uint
	User   User `gorm:"foreignKey:UserId"`
}
