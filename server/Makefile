.PHONY: migrations

# URL to connect to DB for migrations
db_url = "cockroach://$(DB_USER):$(DB_PASSWD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
# title of migration
title = "migration"

# --- #
# APP #
# --- #

dev:
	go run ./cmd/app/main.go

# ------- #
# SWAGGER #
# ------- #

swagger-update:
	@swag_v2 init

swagger-fmt:
	@swag_v2 fmt

swagger: swagger-fmt swagger-update

# ---------- #
# MIGRATIONS #
# ---------- #

check-db-env-vars:
	@if [ -z "$(DB_USER)" ] || [ -z "$(DB_HOST)" ] || [ -z "$(DB_PORT)" ] || [ -z "$(DB_NAME)" ]; then \
		echo "Not all of DB env-vars is presented"; \
		exit 2; \
	fi

# use "title" var for name the migration
migrations: check-db-env-vars
	migrate create -digits 2 -dir migrations -ext sql -seq "$(title)"

migrate-force: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose force 1

migrate-version: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose version

migrate-up: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose up 1

migrate-down: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose down 1
