package db

import (
	"fmt"

	"github.com/dehwyy/Makoto/backend/auth/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgres_user, _     = config.GetOptionByKey("db.auth.postgres.user")
	postgres_password, _ = config.GetOptionByKey("db.auth.postgres.password")
	postgres_port, _     = config.GetOptionByKey("db.auth.postgres.port")
)

type Conn struct {
	DB *gorm.DB
}

func New() *Conn {
	fmt.Println(postgres_user, postgres_password, postgres_port)
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=postgres port=%s sslmode=disable", postgres_user, postgres_password, postgres_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS")
	return &Conn{
		DB: db,
	}
}
