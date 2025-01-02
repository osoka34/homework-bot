DB_HOST ?= localhost
DB_PORT ?= 12000
DB_USER ?= myuser
DB_PASSWORD ?= mypassword
DB_NAME ?= mydb
DB_SSL_MODE ?= disable
MIGRATIONS_DIR ?= migrations

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

MIGRATE ?= migrate

.PHONY: help
help:
	@echo "Использование:"
	@echo "  make migrate-up       - Применить все миграции"
	@echo "  make migrate-down     - Откатить миграции"
	@echo "  make migrate-force    - Принудительно установить версию миграции"
	@echo "  make migrate-status   - Проверить статус миграций"

.PHONY: migrate-up
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down


.PHONY: migrate-force
migrate-force:
ifndef VERSION
	$(error VERSION is required. Use: make migrate-force VERSION=version_number)
endif
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(VERSION)

.PHONY: migrate-status
migrate-status:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version


.PHONY: dep
dep:
	docker-compose up -d


.PHONY: down
down:
	docker-compose down



