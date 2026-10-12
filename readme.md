# Rizat Blog - Service Go API

Backend RESTful API for a blog platform built with **Golang** and **Gin Framework**, following the principles of **Clean Architecture**.

---

## 🛠️ Tech Stack & Package List

Below is the list of core libraries and third-party packages used in this project:

### Core & Framework
* **[Go](https://go.dev/)** `v1.27+` - Core programming language.
* **[Gin Web Framework](https://github.com/gin-gonic/gin)** (`github.com/gin-gonic/gin`) - High-performance HTTP web framework.

### Database & Migration
* **[Postgres](https://postgresql.org)** - Database for account user 
* **[GORM](https://gorm.io/)** (`gorm.io/gorm`) - ORM library for database operations and data seeding.
* **[GORM PostgreSQL Driver](https://gorm.io/driver/postgres)** (`gorm.io/driver/postgres`) - Official PostgreSQL driver for GORM.
* **[golang-migrate](https://github.com/golang-migrate/migrate)** (`github.com/golang-migrate/migrate/v4`) - Database migration tool and library for executing SQL schema migrations (`up`/`down`).

### Security & Utilities
* **[Crypto / Bcrypt](https://golang.org/x/crypto/bcrypt)** (`golang.org/x/crypto/bcrypt`) - Secure password hashing library for user/admin authentication.
* **[godotenv](https://github.com/joho/godotenv)** (`github.com/joho/godotenv`) - Loads environment variables from a `.env` file.
* **[uuid](https://github.com/google/uuid)** (`github.com/google/uuid`) - Generate universal uniq id
* **[Air](https://github.com/air-verse/air)** (`github.com/air-verse/air`) - Live reload tool development

### Code Quality & Security Tools (Dev)
* **[Snyk](https://snyk.io/)** - Static Application Security Testing (SAST) & dependency vulnerability scanner.
Check and scan snyk you can run :
```
snyk auth // for authentication your account
snyk monitor
```

---

## Initial Apps
### Install GoLang
https://go.dev/dl/

### Install packages
```
go mod download
```

## Struktur Folder
```
my-gin-app/
├── cmd/
│   └── api/
│       └── main.go           # Entry point aplikasi
│   └── migrate/
│       └── main.go           # For running migrate
├── database/
│   └── migrations_postgres/  # Folder for migrations
├── internal/
│   ├── config/
│   │   └── environment.go    # Config file environment
│   ├── domain/
│   │   └── health.go         # Model & Interface / Contract
│   ├── service/
│   │   └── health_service.go # Business logic
│   └── handler/
│       └── health_handler.go # HTTP Handler (Gin Context)
├── pkg/
│    └── database/
│        └── seeder/
│            ├── seeder.go          <-- Berisi RunAll(db)
│            └── user_seeder.go     <-- Berisi SeedUsers(db)
│    └── migration.go   # Module for migrate up/down
│    └── postgres.go    # Module connection GORM
├── go.mod
├── go.sum
└── Makefile    # Command for running
```

## Migrations
### Install golang migrate
via homebrew
```
brew install golang-migrate
```
### Create File Migrations
```
migrate create -ext sql -dir database/migrations_postgres -seq name_execute
```

### Run migrations
```
make migrate up
```

### Run migrate down
```
make migrate down
```

### Run seeder
```
make migrate db-seed
```
## Run Apps
### Running apps
```
make run
```

### Running dev
```
make run-dev
```

### Build binary
```
make build
```

### Run production (binary)
```
./app
```
