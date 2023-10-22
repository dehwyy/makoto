package models

type HashmapTag struct {
	Id    uint32 `gorm:"primaryKey"`
	Text  string
	Items []*HashmapItem `gorm:"many2many:hashmap_item_tags;"`
}
