go_entrypoint_path="./cmd/admin_panel/main.go"

# --- #
# APP #
# --- #

dev:
	go run $(go_entrypoint_path)

lint:
	golangci-lint run -c ./.golangci.yml ./...
