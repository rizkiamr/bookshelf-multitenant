run:
	go run ./cmd/api

generate-swagger:
	cd ./cmd/api && swag init

.PHONY: run generate-swagger