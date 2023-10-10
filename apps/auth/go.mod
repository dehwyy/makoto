module github.com/dehwyy/makoto/apps/auth

go 1.21.2

replace (
	github.com/dehwyy/makoto/config => ../../config
	github.com/dehwyy/makoto/libs/grpc => ../../libs/grpc
	github.com/dehwyy/makoto/libs/logger => ../../libs/logger
)

require (
	github.com/dehwyy/makoto/config v0.0.0-00010101000000-000000000000
	github.com/dehwyy/makoto/libs/logger v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.13.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.4
)

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
