.SILENCE:

docker-up:
	(docker network inspect makoto_network || docker network create makoto_network) && docker-compose --env-file .env.prod up --build
