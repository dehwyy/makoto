package models

type HashmapTag struct {
	Id   uint32 `gorm:"primaryKey;"`
	Text string
}
