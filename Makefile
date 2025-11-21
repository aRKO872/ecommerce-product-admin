# Makefile for Docker Compose

# Default service
SERVICE ?= bff-service
MIGRATION_NAME ?= initialize_tables
DB_USER ?= appuser
DB_PASSWORD ?= apppassword
DB_PORT ?= 3306
DB_NAME ?= appdb
DIRTY_VERSION ?= 0
KAFKA_HOST_PORT ?= localhost:9092
KAFKA_TOPIC_NAME ?= post-logs

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

generate-new-migration:
	migrate create -ext sql -dir ./migrations -seq $(MIGRATION_NAME)

migrate-up:
	migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" up

migrate-down:
	migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" down

migrate-down-dirty:
	migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" force $(DIRTY_VERSION)

create-kafka-topic:
	docker exec -it broker /opt/kafka/bin/kafka-topics.sh --create --topic $(KAFKA_TOPIC_NAME) --bootstrap-server $(KAFKA_HOST_PORT)

list-kafka-topics:
	docker exec -it broker /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server $(KAFKA_HOST_PORT)

delete-kafka-topic:
	docker exec -it broker /opt/kafka/bin/kafka-topics.sh --delete --topic $(KAFKA_TOPIC_NAME) --bootstrap-server $(KAFKA_HOST_PORT)