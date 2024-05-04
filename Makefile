dc := docker-compose -f ./docker-compose.yml

up:
	$(dc) up -d --build

down:
	$(dc) down

restart:
	$(dc) restart

reup:
	@make down
	@make up

rm:
	$(dc) down --rmi all

logs:
	$(dc) logs -f

.PHONY: up down restart reup rm logs