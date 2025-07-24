package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	LogLevel string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, reading env from system")
	}

	return &Config{
		AppPort:  getEnv("APP_PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   getEnv("DB_PORT", "5432"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPass:   getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "subscriptions"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
