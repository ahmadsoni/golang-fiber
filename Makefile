include .env

APP_NAME=gofiber_app
MIGRATE_PATH=./db/migrations

.PHONY: up down restart migrate-up migrate-down migrate-version migrate-force

up:
	docker-compose up --build

down:
	docker-compose down

restart:
	docker-compose down && docker-compose up --build

migrate-up:
	docker exec -it gofiber_app migrate -path /app/db/migrations -database "postgres://$$(grep DB_USER .env | cut -d '=' -f2):$$(grep DB_PASSWORD .env | cut -d '=' -f2)@db:5432/$$(grep DB_NAME .env | cut -d '=' -f2)?sslmode=disable" up

migrate-down:
	docker exec -it gofiber_app migrate -path /app/db/migrations -database "postgres://$$(grep DB_USER .env | cut -d '=' -f2):$$(grep DB_PASSWORD .env | cut -d '=' -f2)@db:5432/$$(grep DB_NAME .env | cut -d '=' -f2)?sslmode=disable" down

migrate-version:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable version

migrate-force:
	docker exec -it $(APP_NAME) migrate -path /app/$(MIGRATE_PATH) -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable force

seed-%:
	docker exec -i $(APP_NAME) psql -U $(DB_USER) -d $(DB_NAME) -f /app/db/seeders/$*.sql