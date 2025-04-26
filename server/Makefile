.PHONY: migrations

go_entrypoint_path="./cmd/app/main.go"
server_runner_path="./internal/app/server/server.go"

# URL to connect to DB for migrations
db_url = "cockroach://$(DB_USER):$(DB_PASSWD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
# title of migration
title = "migration"
version = 1

# --- #
# APP #
# --- #

dev:
	go run $(go_entrypoint_path)

lint:
	golangci-lint run -c ./.golangci.yml ./...

clear-logs:
	rm -f ./error.log ./info.log

# ------- #
# SWAGGER #
# ------- #

swagger-update:
	@swag init -g $(server_runner_path)

swagger-fmt:
	@swag fmt -g $(server_runner_path)

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
migrations:
	migrate create -digits 2 -dir migrations -ext sql -seq "$(title)"

# use "version" var for specify the version of the migration to force
migrate-force: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose force $(version)

migrate-version: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose version

migrate-up: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose up 1

migrate-down: check-db-env-vars
	migrate -database $(db_url) -path migrations -verbose down 1
