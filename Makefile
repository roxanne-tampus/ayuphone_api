# Database connection settings
ifneq (,$(wildcard .env))
    include .env
    export
endif
GOOSE=go run github.com/pressly/goose/v3/cmd/goose

# Target to create a new migration
new-migration:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -dir $(MIGRATION_FOLDER) create $$name sql

# Target to apply migrations
migrate-up:
	$(GOOSE) -dir $(MIGRATION_FOLDER) postgres $(DATABASE_URL) up

# Target to rollback last migration
migrate-down:
	$(GOOSE) -dir $(MIGRATION_FOLDER) postgres $(DATABASE_URL) down

# Target to print migration status
migrate-status:
	$(GOOSE) -dir $(MIGRATION_FOLDER) postgres $(DATABASE_URL) status

# Target to fix $(GOOSE) migration version
migrate-fix:
	$(GOOSE) -dir $(MIGRATION_FOLDER) postgres $(DATABASE_URL) fix