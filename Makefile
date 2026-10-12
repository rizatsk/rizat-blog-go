.PHONY: run-dev migrate-up migrate-down db-seed

run-dev:
	go run cmd/api/main.go

migrate-up:
	go run cmd/migrate/main.go -action=up

migrate-down:
	go run cmd/migrate/main.go -action=down

db-seed:
	go run cmd/migrate/main.go -action=seed-up