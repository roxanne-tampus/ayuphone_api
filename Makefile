# Makefile for managing migrations

# Configuration
MIGRATE_PATH=$(HOME)/go/bin/migrate

# Load environment variables
export $(shell grep -v '^#' .env | xargs)

# Commands
migrate-up:
	$(MIGRATE_PATH) -path $(MIGRATE_DIR) -database $(DATABASE_URL) up

migrate-down:
	$(MIGRATE_PATH) -path $(MIGRATE_DIR) -database $(DATABASE_URL) down

migrate-status:
	$(MIGRATE_PATH) -path $(MIGRATE_DIR) -database $(DATABASE_URL) version

# Aliases
up: migrate-up
down: migrate-down
status: migrate-status

.PHONY: migrate-up migrate-down migrate-status up down status
