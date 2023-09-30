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

	res := w.db.Create(&word_model{
		UserId: userId,
		Word:   key,
		Value:  value,
		Extra:  extra,
		Tags:   tags,
	})

	return res.Error
}

func (w *WordsService) RemoveWord(wordId uint) error {

	res := w.schema().Delete(&word_model{}, "id = ?", wordId)

	return res.Error
}

func (w *WordsService) EditWord(wordId uint, new_key, new_value, new_extra string, tags []*tag_model) error {

	res := w.schema().Where("id = ?", wordId).Updates(&word_model{
		Word:  new_key,
		Value: new_value,
		Extra: new_extra,
		Tags:  tags,
	})

	return res.Error
}
