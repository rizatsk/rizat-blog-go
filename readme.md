# Service API With Golang

### Struktur Folder
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

### Migrations
#### Install golang migrate
via homebrew
```
brew install golang-migrate
```
#### Create File Migrations
```
migrate create -ext sql -dir database/migrations_postgres -seq name_execute
```

#### Run migrations
```
make migrate up
```

#### Run migrate down
```
make migrate down
```

#### Run seeder
```
make migrate db-seed
```