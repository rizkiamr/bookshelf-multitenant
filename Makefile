run:
	go run ./cmd/api

generate-swagger:
	cd ./cmd/api && swag init

build:
	cd ./cmd/api && go build -o main main.go

docker-build:
	docker build -t bookshelf-multitenant-api:latest .

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down -v

docker-compose-rebuild:
	docker compose down -v && docker build -t bookshelf-multitenant-api:latest . && docker compose up -d

.PHONY: run generate-swagger docker-build build docker-compose-up docker-compose-down docker-compose-rebuild