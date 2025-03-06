dev:
	go run ./main.go

swag-update:
	@swag_v2 init

swag-fmt:
	@swag_v2 fmt

swagger: swag-fmt swag-update

migrate:
	go run ./main.go migrate
