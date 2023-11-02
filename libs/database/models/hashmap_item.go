package models

import (
	"time"

	"github.com/google/uuid"
)

type HashmapItem struct {
	Id        uint32    `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"index"`
	Key       string
	Value     string
	Extra     string
	CreatedAt time.Time
	Tags      []*HashmapTag `gorm:"many2many:hashmap_item_tags;"`
}
