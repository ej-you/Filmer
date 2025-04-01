go_entrypoint_path="./cmd/app/main.go"

# --- #
# APP #
# --- #

dev:
	go run $(go_entrypoint_path)

lint:
	golangci-lint run -c ./.golangci.yml ./...

clean-lint-cache:
	golangci-lint cache clean

# ------- #
# SWAGGER #
# ------- #
