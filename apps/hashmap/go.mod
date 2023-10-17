module github.com/dehwyy/makoto/apps/hashmap

go 1.21.2

replace (
	github.com/dehwyy/makoto/libs/config => ../../libs/config
	github.com/dehwyy/makoto/libs/database => ../../libs/database
	github.com/dehwyy/makoto/libs/grpc => ../../libs/grpc
	github.com/dehwyy/makoto/libs/logger => ../../libs/logger
	github.com/dehwyy/makoto/libs/middleware => ../../libs/middleware

)

require (
	github.com/dehwyy/makoto/libs/config v0.0.0-00010101000000-000000000000
	github.com/dehwyy/makoto/libs/database v0.0.0-00010101000000-000000000000
	github.com/dehwyy/makoto/libs/grpc v0.0.0-00010101000000-000000000000
	github.com/dehwyy/makoto/libs/logger v0.0.0-20231012081555-72f2af8b2218
	github.com/dehwyy/makoto/libs/middleware v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.3.1
	github.com/twitchtv/twirp v8.1.3+incompatible
	gorm.io/gorm v1.25.5
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gorm.io/driver/postgres v1.5.3 // indirect
)
