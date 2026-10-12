package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port         string
	AppName      string
	AppEnv       string
	DBPGHost     string
	DBPGPort     string
	DBPGUsername string
	DBPGPassword string
	DBPGName     string
	DBPGSSLMode  string
	DBPGSchema   string
}

func LoadEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Println("Info file .env is not found, used default environment variable")
	}

	return &Environment{
		Port:         getEnv("PORT", "8000"),
		AppName:      getEnv("APP_NAME", "rizat-service-blog"),
		AppEnv:       getEnv("APP_ENV", "development"),
		DBPGHost:     getEnv("DB_PG_HOST", "localhost"),
		DBPGPort:     getEnv("DB_PG_PORT", "5432"),
		DBPGUsername: getEnv("DB_PG_USERNAME", "postgres"),
		DBPGPassword: getEnv("DB_PG_PASSWORD", ""),
		DBPGName:     getEnv("DB_PG_NAME", "my_db"),
		DBPGSSLMode:  getEnv("DB_PG_SSLMODE", "disable"),
		DBPGSchema:   getEnv("DB_PG_SCHEMA", "public"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
