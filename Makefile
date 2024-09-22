# Variables
GOOSE_CMD = goose
DB_DRIVER = postgres # Adjust this to your database driver, e.g., postgres, mysql, sqlite
DB_DSN = "user=youruser password=yourpass dbname=yourdb sslmode=disable" # Adjust with your connection string
MIGRATION_DIR = migrations

# Help command (optional)
.PHONY: help
help:
	@echo "Available migration commands:"
	@echo "  make migration new          - Create a new migration"
	@echo "  make migration status       - Get migration status"
	@echo "  make migration up           - Run all migrations"
	@echo "  make migration down         - Rollback the most recent migration"
	@echo "  make migration reset        - Rollback all migrations and reapply them"

# Migration group
.PHONY: migration
migration: help

# Create a new migration
.PHONY: migration-new
migration-new:
	@read -p "Enter migration name: " name; \
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) create $$name sql

# Check migration status
.PHONY: migration-status
migration-status:
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) status

# Migrate up
.PHONY: migration-up
migration-up:
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) up

# Migrate down
.PHONY: migration-down
migration-down:
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) down

# Reset all migrations (down then up)
.PHONY: migration-reset
migration-reset:
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) reset
