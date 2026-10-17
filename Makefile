.PHONY: build run run-dev migrate-up migrate-down migrate-down-step db-seed

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

migrate-down-step:
	go run cmd/migrate/main.go -action=down-step

db-seed:
	go run cmd/migrate/main.go -action=seed-up