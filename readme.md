# Template - Service Go API With GIN

Backend RESTful API for a platform built with **Golang** and **Gin Framework**, following the principles of **Clean Architecture**.

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
* **[JWT](https://github.com/golang-jwt/jwt/v5)** (`github.com/golang-jwt/jwt/v5`) - JWT for token auth
* **[Validator](https://github.com/go-playground/validator/v10)** (`github.com/go-playground/validator/v10`) - Validator request

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

go mod tidy
```

## Struktur Folder
```
my-gin-app/
├── cmd/
│   └── api/
│       └── main.go                         # Entry point aplikasi
│   └── migrate/
│       └── main.go                         # For running migrate
├── database/
│   └── migrations_postgres/                # Folder for migrations
├── internal/
│   ├── config/
│   │   └── environment.go                  # Config file environment
│   ├── domain/
│   │   └── health.go                       # Model & Interface / Contract
│   ├── repository/
│   │   └── user_repository.go              # Logic for get data contract from domain
│   ├── service/
│   │   └── health_service.go               # Business logic
│   └── handler/
│       ├── middleware/
│       │   ├── auth_middleware.go          # Checking valid token
│       │   └── refreshtoken_middleware.go  # Checking valid refresh token
│       └── health_handler.go               # HTTP Handler (Gin Context)
├── pkg/
│    ├── asynclock/
│    |   └── sync_lock.go                   # For snyc loc with mutex  
│    ├── database/
│    │   ├── seeder/
│    │   │  ├── seeder.go                   <-- Berisi RunAll(db)
│    │   │  └── user_seeder.go              <-- Berisi SeedUsers(db)
│    │   ├── migration.go                   # Module for migrate up/down
│    │   └── postgres.go                    # Module connection GORM
│    ├── response/
│    │   ├── new_error.go                   # Send Error function to handler
│    │   ├── response_api.go                # Handle error send to response
│    │   └── validator.go                   # Handle validator from error validation json
│    ├── time/
│    │   └── time.go                        # Handle get time now used format
│    └── tokenize/
│        └── tokenize.go  
├── .ait.toml                               # For run mode dev with Air
├── .env.example
├── .gitignor
├── go.mod
├── go.sum
└── Makefile                    # Command for running
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
make migrate-up
```

### Run migrate down
```
make migrate-down
```

### Run migrate down per step
```
make migrate-down-step
```

### Run seeder
```
make db-seed
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

## Implement Goroutines and Channel
Lock-nya sendiri udah benar. Masalahnya, **function `RefreshToken` return duluan sebelum goroutine-nya selesai**, jadi `defer Unlock` jalan lebih awal dan lock dilepas waktu DELETE/INSERT masih jalan di belakang.

### Biang keroknya

[auth_service.go:173](internal/service/auth_service.go#L173):

```go
go func() {
    waitGroup.Wait()   // yang nunggu cuma goroutine ini
    close(errChan)
}()

if len(errChan) > 0 { // dicek langsung, nggak nunggu apa-apa
```

`waitGroup.Wait()` ada di dalam goroutine, jadi yang ngeblok cuma goroutine itu, bukan function utama. Function utama langsung ngecek `len(errChan)`. Di titik itu DELETE dan INSERT hampir pasti belum selesai, jadi channel masih kosong, `len == 0`, dan function langsung return sukses.

### Urutan kejadiannya

```
Request A (user X)                          Request B (user X, token sama)
──────────────────                          ──────────────────────────────
Lock("X")            ✅ dapat lock
GetAuthTokenById     ✅ token ada
generate token
spawn goroutine DELETE ─┐
spawn goroutine INSERT ─┤ (masih jalan di background)
spawn goroutine Wait   ─┤
len(errChan) == 0       │
return sukses           │
defer Unlock("X") 🔓    │                   Lock("X")  ✅ dapat lock (A udah unlock)
                        │                   GetAuthTokenById ✅ token MASIH ADA
                        │                     (DELETE punya A belum selesai)
                   DELETE selesai           ...B juga sukses refresh 💥
```

Hasilnya, satu refresh token bisa dipakai dua kali. Lock-nya cuma melindungi bagian sebelum goroutine di-spawn. Kerjaan di goroutine udah di luar lock.

Ada efek samping lain juga:
- **Error nggak pernah ketangkep.** Kalau DELETE atau INSERT gagal, client tetap dapat token baru, padahal INSERT-nya mungkin gagal. Akibatnya token barunya nggak valid pas dipakai nanti.
- **Datanya bisa setengah jadi.** DELETE sukses tapi INSERT gagal (atau sebaliknya) nggak ada rollback-nya.
- Lo terpaksa pakai `context.Background()` mungkin karena context request udah ke-cancel duluan, soalnya handler udah return. Ini gejala dari masalah yang sama.

### Cara kerja masing-masing

**1. Locking (`SyncLock`, [asynclock.go](pkg/asynclock/asynclock.go))**
- `mu` itu mutex global yang cuma ngejaga map `locks` dan `counts`. Dia dipegang sebentar aja.
- Tiap key (userId) punya `*sync.Mutex` sendiri. `Lock(key)` naikin `counts`, lepas `mu`, lalu nunggu mutex milik key itu. Jadi request user X nggak ngeblok user Y.
- `Unlock` turunin count. Kalau count-nya 0, entry-nya dihapus dari map biar memori nggak bocor.
- Implementasi ini udah oke. Aturannya: lock cuma melindungi kode yang **jalan di antara `Lock` dan `Unlock`, di goroutine yang sama alurnya**. Kalau kerjaannya dilempar ke goroutine lain lalu function return, lock nggak ikut nungguin.

**2. Goroutine**
- `go func(){...}()` artinya "jalanin ini secara paralel, gw nggak nunggu". Function pemanggil langsung lanjut ke baris berikutnya.
- Goroutine nggak terikat sama lifetime function yang nge-spawn. Dia tetap jalan walaupun function-nya udah return.

**3. WaitGroup**
- `Add(n)` set counter, `Done()` ngurangin counter, `Wait()` **ngeblok goroutine yang manggil** sampai counter 0.
- Kuncinya ada di siapa yang manggil `Wait()`. Kalau mau function utama yang nunggu, `Wait()` harus dipanggil langsung di function utama.

**4. Channel**
- `make(chan error, 2)` itu buffered channel. Kirim ke channel ini nggak ngeblok selama buffer belum penuh.
- `len(ch)` cuma snapshot isi buffer **saat itu juga**. Dia nggak nunggu apa-apa, jadi nggak bisa dipakai buat sinkronisasi.
- Yang ngeblok itu `<-ch` atau `for err := range ch`. Range berhenti setelah channel di-`close`.

### Perbaikan

**Opsi A, minimal: tunggu dulu sebelum cek error**

```go
// Delete & insert goroutine tetap sama ...

waitGroup.Wait() // dipanggil di function utama, jadi function ngeblok di sini
close(errChan)

for err := range errChan {
    if err != nil {
        return nil, response.NewAppError(500, "", nil)
    }
}

return &domain.Login{TokenString: tokenString, RefreshTokenString: refreshTokenString}, nil
```

Dengan ini `Unlock` baru jalan setelah DELETE dan INSERT selesai, dan lo bisa balik pakai `ctx.Request.Context()`, nggak perlu `context.Background()` lagi.

**Opsi B, yang gw saranin: nggak usah goroutine**

DELETE dan INSERT itu dua query kecil. Kalau diparalelin, paling hemat beberapa milidetik, tapi lo kehilangan urutan dan atomicity. Untuk rotasi refresh token, idealnya keduanya dalam **satu DB transaction**:

```go
// sequential, di dalam lock
if err := authService.authRepo.DeleteTokenById(ctx.Request.Context(), dataToken.AuthId); err != nil {
    log.Printf("Failed delete token by id : %v", err)
    return nil, response.NewAppError(500, "", nil)
}
if err := authService.authRepo.InsertToken(ctx.Request.Context(), authTokenId, user.ID); err != nil {
    log.Printf("Failed create token from table : %v", err)
    return nil, response.NewAppError(500, "", nil)
}
```

Lebih bagus lagi kalau repo punya method `RotateToken(ctx, oldId, newId, userId)` yang jalanin DELETE + INSERT dalam satu transaksi. Cek "token masih ada" juga bisa digabung lewat `DELETE ... WHERE id = ?` lalu dicek `RowsAffected == 1`. Dengan begitu proteksi reuse-nya tetap jalan walaupun app lo jalan di beberapa instance, karena `SyncLock` cuma berlaku di dalam satu proses.

Mau gw terapin Opsi A atau B ke file-nya?