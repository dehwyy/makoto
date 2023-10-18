package repository

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemsRepository struct {
	l  logger.Logger
	db *gorm.DB
}

func NewItemsRepository(db *gorm.DB, l logger.Logger) *ItemsRepository {
	return &ItemsRepository{
		l:  l,
		db: db,
	}
}

func (i *ItemsRepository) schema() *gorm.DB {
	return i.db.Model(&models.HashmapItem{})
}

func (i *ItemsRepository) GetItems(userId uuid.UUID) (items []*models.HashmapItem, err error) {
	return items, i.schema().Preload("Tags").Where("user_id = ?", userId).Find(&items).Error
}

func (i *ItemsRepository) CreateItem(userId uuid.UUID, key, value, extra string, tags []*models.HashmapTag) error {

	new_item := &models.HashmapItem{
		UserId: userId,
		Key:    key,
		Value:  value,
		Extra:  extra,
	}

	// Create `Word` and add to it `Tags`
	// ? Transaction

	return i.db.Transaction(func(tx *gorm.DB) error {
		if err := i.db.Create(new_item).Error; err != nil {
			return err
		}

		return i.db.Model(new_item).Association("Tags").Append(tags)
	})
}

func (i *ItemsRepository) RemoveItem(user_id uuid.UUID, ItemId uint32) error {

	// Removing data from `Item` and its tags
	// ? Transaction
	return i.db.Transaction(func(tx *gorm.DB) error {

		err := i.db.Model(&models.HashmapItem{
			Id:     ItemId,
			UserId: user_id,
		}).Association("Tags").Clear()

		if err != nil {
			return err
		}

		return i.schema().Delete("id = ?", ItemId).Error
	})
}

func (i *ItemsRepository) EditItem(user_id uuid.UUID, itemId uint32, new_key, new_value, new_extra string, tags []*models.HashmapTag) error {

	payload := &models.HashmapItem{
		Key:   new_key,
		Value: new_value,
		Extra: new_extra,
	}

	// Update `Item` and change `Tags`
	// ? Transaction
	return i.db.Transaction(func(tx *gorm.DB) error {

		if err := i.schema().Where("id = ? and user_id = ?", itemId).Updates(payload).Error; err != nil {
			return err
		}

		return i.db.Model(&models.HashmapItem{
			Id: itemId,
		}).Association("Tags").Replace(tags)
	})
}
