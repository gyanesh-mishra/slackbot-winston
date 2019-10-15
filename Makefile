#!make

run:
	@echo "+\n++ Running application ...\n+"
	@docker-compose -f deployments/docker-compose.yml  up

images:
	@echo "+\n++ Building images ...\n+"
	@docker-compose -f deployments/docker-compose.yml  build --parallel

stop:
	@echo "+\n++ Stopping application ...\n+"
	@docker-compose -f deployments/docker-compose.yml  down -t 2

clean:
	@echo "+\n++ Removing containers, images, volumes etc...\n+"
	@docker-compose -f deployments/docker-compose.yml  down --rmi all --volumes
	@docker-compose -f deployments/docker-compose.yml  rm -f -v -s
