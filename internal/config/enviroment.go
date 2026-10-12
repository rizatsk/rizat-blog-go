package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	AppName string
	AppEnv  string
}

func LoadEnvironment() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Info file .env is not found, used default environment variable")
	}

	return &Config{
		Port:    getEnv("PORT", "8000"),
		AppName: getEnv("APP_NAME", "rizat-service-blog"),
		AppEnv:  getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
