package service

import (
	database "github.com/dehwyy/makoto/backend/dict/db"
	"github.com/dehwyy/makoto/backend/dict/db/models"
	"github.com/dehwyy/makoto/backend/dict/logger"
	"gorm.io/gorm"
)

type TagsService struct {
	l  logger.AppLogger
	db *gorm.DB
}

type tag_model = models.Tag

func NewTagsService(l logger.AppLogger, connection *database.Conn) *TagsService {
	return &TagsService{
		l:  l,
		db: connection.DB,
	}
}

func (t *TagsService) GetTagOrCreate(text string) *tag_model {
	var tag *tag_model

	// trying to get `tag` from DB
	// if it doesn't exist -> error would appear -> new `tag` record would be created
	tag, err := t.GetTag(text)

	if err != nil {
		tag = t.CreateTag(text)
	}

	return tag
}

func (t *TagsService) GetAllTags() []tag_model {
	var tags []tag_model

	t.db.Model(&tag_model{}).Find(&tags)

	return tags
}

func (t *TagsService) GetTag(text string) (*tag_model, error) {

	// struct to store value
	var tag *tag_model

	res := t.db.Model(&tag_model{}).Where("text = ?", text).First(&tag)

	// if I'm not mistaking, res.Error would occur only when the `tag` wasn't found
	return tag, res.Error
}

func (t *TagsService) CreateTag(text string) *tag_model {

	tag := &tag_model{
		Text: text,
	}

	// `res` is not neccessary
	t.db.Create(tag)

	// no error returns
	return tag
}

func (t *TagsService) TagsFromStringArray(string_tags []string) []*tag_model {
	var tags []*models.Tag
	for _, tag := range string_tags {
		db_tag := t.GetTagOrCreate(tag)
		tags = append(tags, db_tag)
	}

	return tags
}
