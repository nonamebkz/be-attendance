include .env

# Makefile for generating Swagger documentation

# Path to the swag binary
SWAG_BIN=/home/noname/go/bin/swag



# Target to generate Swagger documentation
swagger:
	$(SWAG_BIN) init

# Migration commands
create_migration:
	migrate create -ext sql -dir database/migration -seq $(name)

migrate_up:
	migrate -path database/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose  up

migrate_down:
	migrate -path database/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose  down

force_migration:
	migrate -path database/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $(version)

# Optional: Clean target to remove generated files
clean:
	rm -rf docs

# Default target
.PHONY: swagger clean create_migration migrate_up migrate_down force_migration