#!make

docker-compose = docker-compose -f deployments/docker-compose-dev.yml
docker-compose-prod = docker-compose -f deployments/docker-compose-prod.yml

app:
	@echo "+\n++ Running application in background...\n+"
	@$(docker-compose) up

database:
	@echo "+\n++ Running database in background...\n+"
	@$(docker-compose) up -d winston-database winston-database-gui

build-local:
	@echo "+\n++ Building images for local development ...\n+"
	@$(docker-compose) build --parallel

build-prod:
	@echo "+\n++ Building images for production use...\n+"
	@$(docker-compose-prod) build --parallel

app-prod:
	@echo "+\n++ Running application in production mode...\n+"
	@$(docker-compose-prod) up

stop:
	@echo "+\n++ Stopping application ...\n+"
	@$(docker-compose) down -t 2

clean:
	@echo "+\n++ Removing containers, images, volumes etc...\n+"
	@$(docker-compose) down --rmi all --volumes
	@$(docker-compose) rm -f -v -s
