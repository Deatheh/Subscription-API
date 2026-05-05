COMPOSE_FILE := docker-compose.yaml
ENV_FILE := .env

ifneq (,$(wildcard $(ENV_FILE)))
    include $(ENV_FILE)
    export
endif

MIGRATIONS_PATH := /migrations
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable


.PHONY: up
up: 
	docker compose -f $(COMPOSE_FILE) up -d

.PHONY: down
down: 
	docker compose -f $(COMPOSE_FILE) down

.PHONY: reset
reset:
	docker compose -f $(COMPOSE_FILE) down -v

.PHONY: logs
logs: 
	docker compose -f $(COMPOSE_FILE) logs -f

.PHONY: build
build: 
	docker compose -f $(COMPOSE_FILE) build app


.PHONY: migrate-create
migrate-create: ## Создать новые файлы миграции. Использовать: make migrate-create name=название_миграции
	@if [ -z "$(name)" ]; then \
		echo "Ошибка: укажите имя миграции: make migrate-create name=my_migration"; \
		exit 1; \
	fi
	docker run --rm -v ./internal/repository/db/migrations:/migrations migrate/migrate:v4.18.1 create -ext sql -dir /migrations -seq $(name)

.PHONY: migrate-up
migrate-up: 
	docker run --rm --network subscription_default -v ./internal/repository/db/migrations:/migrations migrate/migrate:v4.18.1 -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: 
	docker run --rm --network subscription_default -v ./internal/repository/db/migrations:/migrations migrate/migrate:v4.18.1 -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

.PHONY: migrate-reset
migrate-reset:
	docker run --rm --network subscription_default -v ./internal/repository/db/migrations:/migrations migrate/migrate:v4.18.1 -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down -all
