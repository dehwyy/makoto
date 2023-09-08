.SILENT:

name = auth

gen-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./grpc/auth.proto

docker-clear:
	docker rmi -f $(name)-${name}

docker-up: docker-clear
	docker-compose up --build
