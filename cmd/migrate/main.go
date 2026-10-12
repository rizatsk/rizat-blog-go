package main

import (
	"flag"
	"log"
	"service_blog_go/internal/config"
	"service_blog_go/pkg/database"
	"service_blog_go/pkg/database/seeder"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	action := flag.String("action", "up", "Action command: 'up', 'down', or 'seed-up'")
	flag.Parse()

	env := config.LoadEnvironment()

	switch *action {
	case "up":
		log.Println("Running migrate up...")
		database.RunMigrationsUp(env)
	case "down":
		log.Println("Running migrate down...")
		database.RunMigrationsDown(env)
	case "seed-up":
		db_postgres := database.InitPostgres(env)
		seeder.RunAll(db_postgres, env)
	}
}
