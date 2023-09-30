package database

import (
	"fmt"

	"github.com/dehwyy/makoto/backend/dict/config"
	"github.com/dehwyy/makoto/backend/dict/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	DB *gorm.DB
}

var (
	postgres_user     = config.GetOptionByKey("db.postgres.user")
	postgres_password = config.GetOptionByKey("db.postgres.password")
	postgres_port     = config.GetOptionByKey("db.postgres.port")
	dsn               = fmt.Sprintf("host=db user=%s password=%s dbname=postgres port=%s sslmode=disable", postgres_user, postgres_password, postgres_port)
)

func New() *Conn {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	return &Conn{
		DB: db,
	}
}

func (c *Conn) RunAllMigrations() {
	if err := c.DB.AutoMigrate(&models.Tag{}, &models.Word{}); err != nil {
		panic(err)
	}
}
