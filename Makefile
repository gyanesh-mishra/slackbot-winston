#!make

docker-compose = docker-compose -f deployments/docker-compose.yml

rerun : images | run

local-dev:
	@echo "+\n++ Running application locally for development ...\n+"
	if ! test -s air; then curl -fLo air https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air && chmod +x air; fi
	./air -c air.conf
	endif

run:
	@echo "+\n++ Running application ...\n+"
	@$(docker-compose) up

database:
	@echo "+\n++ Running database in daemon ...\n+"
	@$(docker-compose) up -d winston-database

images:
	@echo "+\n++ Building images ...\n+"
	@$(docker-compose) build --parallel

stop:
	@echo "+\n++ Stopping application ...\n+"
	@$(docker-compose) down -t 2

clean:
	@echo "+\n++ Removing containers, images, volumes etc...\n+"
	@$(docker-compose) down --rmi all --volumes
	@$(docker-compose) rm -f -v -s
