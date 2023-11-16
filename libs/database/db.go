package database

import (
	"time"

	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(database_url string, l logger.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(database_url), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	d, err := db.DB()
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(10)
	d.SetConnMaxIdleTime(1 * time.Minute)
	// run migrations
	err = db.AutoMigrate(&models.UserData{}, &models.UserToken{}, &models.HashmapTag{},
		&models.HashmapItem{}, &models.Language{}, &models.UserInfo{})
	if err != nil {
		l.Fatalf("Cannot run migrations!")
	}

	l.Infof("Succesfully run db migrate and open connection!")

	// insert into Language table if empty
	var language_records int64
	db.Model(&models.Language{}).Count(&language_records)
	if language_records == 0 {
		l.Infof("Inserting default languages into database...")
		languages := []models.Language{
			{ID: 1, Lang: "english"},
			{ID: 2, Lang: "russian"},
			{ID: 3, Lang: "japanese"},
		}

		err = db.CreateInBatches(languages, len(languages)).Error
		if err != nil {
			l.Fatalf("Cannot insert default languages into database!")
		}
	}

	return db
}
