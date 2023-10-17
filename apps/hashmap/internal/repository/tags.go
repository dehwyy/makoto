package repository

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"gorm.io/gorm"
)

type TagsRepository struct {
	l  logger.Logger
	db *gorm.DB
}

func NewTagsRepository(db *gorm.DB, l logger.Logger) *TagsRepository {
	return &TagsRepository{
		l:  l,
		db: db,
	}
}

func (t *TagsRepository) GetTagOrCreate(text string) (tag *models.HashmapTag) {
	// trying to get `tag` from DB
	// if it doesn't exist -> error would appear -> new `tag` record would be created
	tag, err := t.GetTag(text)

	if err != nil {
		tag = t.CreateTag(text)
	}

	return tag
}

func (t *TagsRepository) GetAllTags() (tags []models.HashmapTag, err error) {
	return tags, t.db.Model(&models.HashmapTag{}).Find(&tags).Error
}

func (t *TagsRepository) GetTag(text string) (tag *models.HashmapTag, err error) {
	return tag, t.db.Model(&models.HashmapTag{}).Where("text = ?", text).First(&tag).Error
}

func (t *TagsRepository) CreateTag(text string) *models.HashmapTag {

	tag := &models.HashmapTag{
		Text: text,
	}

	// `res` is not neccessary
	t.db.Create(tag)

	// no error returns
	return tag
}

func (t *TagsRepository) TagsFromStringArray(string_tags []string) (tags []*models.HashmapTag) {
	for _, tag := range string_tags {
		db_tag := t.GetTagOrCreate(tag)
		tags = append(tags, db_tag)
	}

	return tags
}
