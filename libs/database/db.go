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
	db.AutoMigrate(&models.UserData{}, &models.UserToken{}, &models.HashmapTag{}, &models.HashmapItem{})

	return db
}
