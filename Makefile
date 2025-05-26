include .env

.PHONY: \
	up down restart \
	migrate-up migrate-down migrate-version migrate-force migrate-new soft-migrate \
	docker-migrate-up docker-migrate-down docker-migrate-version docker-migrate-force \
	seed-%

APP_NAME = gofiber_app
MIGRATE_PATH = ./db/migrations

# -----------------------------------
# Docker Compose Commands
# -----------------------------------

up:
	docker-compose up --build

down:
	docker-compose down

restart:
	docker-compose down && docker-compose up --build

# -----------------------------------
# Local Migration Commands
# -----------------------------------

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migrate-version:
	migrate -path $(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

migrate-force:
	migrate -path $(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force

migrate-new:
	@read -p "Migration name: " name; \
	timestamp=$$(date +%s); \
	touch "$(MIGRATE_PATH)/$${timestamp}_$$name.up.sql" "$(MIGRATE_PATH)/$${timestamp}_$$name.down.sql"

# -----------------------------------
# Docker Migration Commands
# -----------------------------------

docker-migrate-up:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_NAME)?sslmode=disable" up

docker-migrate-down:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_NAME)?sslmode=disable" down

docker-migrate-version:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_NAME)?sslmode=disable" version

docker-migrate-force:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_NAME)?sslmode=disable" force

# -----------------------------------
# Seeder (Docker Only)
# -----------------------------------

seed-%:
	docker exec -i $(APP_NAME) psql -U $(DB_USER) -d $(DB_NAME) -f /app/db/seeders/$*.sql
