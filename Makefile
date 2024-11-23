include .env
export $(shell sed 's/=.*//' .env)

# Variables
GO := go
MIGRATE := sqlite3
DB_PATH := internal/db/app.db
MIGRATIONS_PATH := internal/db/migrations

# Commands
.PHONY: all migrate-up migrate-down migrate-prod-up migrate-prod-down build run 

all: migrate build

migrate-up:
	@echo "Running database migrations..."
	$(MIGRATE) $(DB_PATH) < $(MIGRATIONS_PATH)/001_create_initial_schema_up.sql
	$(MIGRATE) $(DB_PATH) < $(MIGRATIONS_PATH)/seeder/seeder.sql

migrate-down:
	@echo "Dropping all tables..."
	$(MIGRATE) $(DB_PATH) < $(MIGRATIONS_PATH)/001_create_initial_schema_down.sql

migrate-prod-up:
	@echo "Running Turso database migrations..."
	$(GO) run $(MIGRATIONS_PATH)/migrate.go --operation=up

migrate-prod-down:
	@echo "Dropping all tables..."
	$(GO) run $(MIGRATIONS_PATH)/migrate.go --operation=down

build:
	@echo "Building Go application..."
	$(GO) build -o main cmd/lambda/main.go

run:
	@echo "Running Go application..."
	$(GO) run cmd/main.go


