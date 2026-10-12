.PHONY: build run run-dev migrate-up migrate-down db-seed

build:
	go build -o app cmd/api/main.go

run:
	go run cmd/api/main.go

run-dev:
	air

migrate-up:
	go run cmd/migrate/main.go -action=up

migrate-down:
	go run cmd/migrate/main.go -action=down

db-seed:
	go run cmd/migrate/main.go -action=seed-up