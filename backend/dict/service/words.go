package service

import (
	database "github.com/dehwyy/makoto/backend/dict/db"
	"github.com/dehwyy/makoto/backend/dict/db/models"
	"github.com/dehwyy/makoto/backend/dict/logger"
	"gorm.io/gorm"
)

type WordsService struct {
	l  logger.AppLogger
	db *gorm.DB
}

type word_model = models.Word

func NewWordsService(l logger.AppLogger, connection *database.Conn) *WordsService {
	return &WordsService{
		l:  l,
		db: connection.DB,
	}
}

func (w *WordsService) schema() *gorm.DB {
	return w.db.Model(&word_model{})
}

func (w *WordsService) GetWords(userId string) ([]*word_model, error) {
	var words []*word_model
	res := w.schema().Preload("Tags").Where("user_id = ?", userId).Find(&words)

	return words, res.Error
}

func (w *WordsService) CreateWord(userId, key, value, extra string, tags []*tag_model) error {

	new_word := &word_model{
		UserId: userId,
		Word:   key,
		Value:  value,
		Extra:  extra,
	}

	// Create `Word` and add to it `Tags`
	// ? Transaction
	err := w.db.Transaction(func(tx *gorm.DB) error {
		create_word := w.db.Create(new_word)

		if create_word.Error != nil {
			return create_word.Error
		}

		add_tags_err := w.db.Model(new_word).Association("Tags").Append(tags)

		return add_tags_err
	})

	return err
}

func (w *WordsService) RemoveWord(wordId uint32) error {

	// Removing data from `Word` and its tags
	// ? Transaction
	err := w.db.Transaction(func(tx *gorm.DB) error {

		err := w.db.Model(&word_model{
			Id: wordId,
		}).Association("Tags").Clear()

		if err != nil {
			return err
		}

		res := w.schema().Delete(&word_model{}, "id = ?", wordId)

		return res.Error
	})

	return err
}

func (w *WordsService) EditWord(wordId uint32, new_key, new_value, new_extra string, tags []*tag_model) error {

	payload := &word_model{
		Word:  new_key,
		Value: new_value,
		Extra: new_extra,
	}

	// Update `Word` and change `Tags`
	// ? Transaction
	err := w.db.Transaction(func(tx *gorm.DB) error {

		update_word := w.schema().Where("id = ?", wordId).Updates(payload)

		if update_word.Error != nil {
			return update_word.Error
		}

		err := w.db.Model(&word_model{
			Id: wordId,
		}).Association("Tags").Replace(tags)

		return err
	})

	return err
}
