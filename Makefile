# Makefile for Docker Compose

# Default service
SERVICE ?= bff-service

up:
	docker compose up -d

up-logs:
	docker compose up

down:
	docker compose down

build:
	docker compose build

build-service:
	docker compose build $(SERVICE)

rebuild-service:
	docker compose rm -sf $(SERVICE) && docker rmi ecommerce-product-admin-$(SERVICE) && docker compose build $(SERVICE) && docker compose up -d $(SERVICE)

restart:
	docker compose down && docker compose up -d

restart-service:
	docker compose stop $(SERVICE) && docker compose rm -f $(SERVICE) && docker compose up -d $(SERVICE)

logs:
	docker compose logs -f $(SERVICE)

ps:
	docker ps

ps-all:
	docker ps -a

clean:
	docker compose down --volumes --rmi all
