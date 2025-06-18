go_entrypoint_path="./cmd/movie_sync/main.go"

# --- #
# APP #
# --- #

dev:
	go run $(go_entrypoint_path)

lint:
	golangci-lint run -c ./.golangci.yml ./...
