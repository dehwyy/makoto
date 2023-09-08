.SILENT:

run-distr:
	cd distributor && make docker-up

run-auth:
	cd auth && make docker-up

docker-create-network:
	docker network create makoto_micro

setup:
	make dockker-create-network
