package main

import (
	"os"

	makoto_config "github.com/dehwyy/makoto/config/src/go"
	makoto_db "github.com/dehwyy/makoto/libs/db/src/go"
	makoto_log "github.com/dehwyy/makoto/libs/logger/src"
)

func main() {
	var db_url string

	if len(os.Args) > 1 {
		db_url = os.Args[1]
	} else {
		db_url = makoto_config.NewConfig().DatabaseDsn
	}

	makoto_log.New().Infof("db_url: %s", db_url)

	makoto_db.New(db_url)
}
