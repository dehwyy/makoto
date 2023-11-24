package main

import (
	"github.com/dehwyy/makoto/libs/database"
	"github.com/dehwyy/makoto/libs/logger"
)

func main() {
	db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", logger.New())

	db.Exec("DELETE FROM user_tokens")
	db.Exec("DELETE FROM user_data")
	db.Exec("DELETE FROM users_languages")
	db.Exec("DELETE FROM user_infos")
}
