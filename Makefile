.SILENCE:

docker-up:
	docker-compose --env-file .env.prod up --build
