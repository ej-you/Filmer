.PHONY: migrations

go_server_path="./cmd/server/main.go"
go_migrator_path="./cmd/migrator/main.go"
server_runner_path="./internal/app/server/server.go"

# title of migration
title = "migration"
version = 1

# --- #
# APP #
# --- #

dev:
	go run $(go_server_path)

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

# use "title" var for name the migration
migrations:
	migrate create -digits 2 -dir migrations -ext sql -seq "$(title)"

# use "version" var for specify the version of the migration to force
migrate-force:
	@go run $(go_migrator_path) force $(version)

migrate-status:
	@go run $(go_migrator_path) status

migrate-up:
	@go run $(go_migrator_path) up -n 1

migrate-down:
	@go run $(go_migrator_path) down -n 1
