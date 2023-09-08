.SILENT:""

name = distributor

gql-gen:
	go run github.com/99designs/gqlgen generate

docker-clear:
	docker rmi -f $(name)-${name}

docker-up: docker-clear
	docker-compose up --build
