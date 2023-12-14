package makoto_db

import (
	"github.com/dehwyy/makoto/libs/db/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(database_url string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: database_url,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db_settings, err := db.DB()
	if err != nil {
		panic(err)
	}

	db_settings.SetMaxOpenConns(100)
	db_settings.SetMaxIdleConns(10)
	db_settings.SetConnMaxIdleTime(5)

	err = db.AutoMigrate(models.UserCredentials{})
	if err != nil {
		panic(err)
	}

	return db
}
