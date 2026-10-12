### Struktur Folder
```
my-gin-app/
├── cmd/
│   └── api/
│       └── main.go           # Entry point aplikasi
├── internal/
│   ├── domain/
│   │   └── health.go         # Model & Interface / Contract
│   ├── service/
│   │   └── health_service.go # Business logic
│   └── handler/
│       └── health_handler.go # HTTP Handler (Gin Context)
├── pkg/
|    └── database/
|    └── postgres.go # Module connection GORM
├── go.mod
└── go.sum
```