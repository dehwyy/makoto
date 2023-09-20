package database

import (
	"fmt"

	"github.com/dehwyy/Makoto/backend/auth/config"
	"github.com/dehwyy/Makoto/backend/auth/db/models"
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgres_user     = config.GetOptionByKey("db.postgres.user")
	postgres_password = config.GetOptionByKey("db.postgres.password")
	postgres_port     = config.GetOptionByKey("db.postgres.port")
	dsn               = fmt.Sprintf("host=db user=%s password=%s dbname=postgres port=%s sslmode=disable", postgres_user, postgres_password, postgres_port)
)

type Conn struct {
	DB *gorm.DB
	l  logger.AppLogger
}

func New(l logger.AppLogger) *Conn {
	// open db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		l.Fatalf("failed to connect database: %v", err)
	}

	return &Conn{
		DB: db,
		l:  l,
	}
}

func (c *Conn) RunAllMigrations() {
	if err := c.DB.AutoMigrate(&models.Token{}, &models.Credentials{}); err != nil {
		c.l.Fatalf("failed to run migrations: %v", err)
	}
}

func (c *Conn) Transaction(fn func(tx *gorm.DB) error) error {
	return c.DB.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
