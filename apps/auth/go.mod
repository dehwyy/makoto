module github.com/dehwyy/makoto/apps/auth

go 1.21.2

replace (
	github.com/dehwyy/makoto/libs/config => ../../libs/config
	github.com/dehwyy/makoto/libs/database => ../../libs/database
	github.com/dehwyy/makoto/libs/grpc => ../../libs/grpc
	github.com/dehwyy/makoto/libs/logger => ../../libs/logger
)

require (
	github.com/dehwyy/makoto/libs/config v0.0.0-00010101000000-000000000000
	github.com/dehwyy/makoto/libs/database v0.0.0-20231016182119-ee1c9c5c6502
	github.com/dehwyy/makoto/libs/grpc v0.0.0-20231016182119-ee1c9c5c6502
	github.com/dehwyy/makoto/libs/logger v0.0.0-20231016182119-ee1c9c5c6502
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/google/uuid v1.4.0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/ravener/discord-oauth2 v0.0.0-20230514095040-ae65713199b3
	github.com/twitchtv/twirp v8.1.3+incompatible
	golang.org/x/crypto v0.14.0
	golang.org/x/oauth2 v0.13.0
	gorm.io/gorm v1.25.5
)

require (
	cloud.google.com/go/compute v1.23.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gorm.io/driver/postgres v1.5.3 // indirect
)
