package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Host        string
	DBDriver    string
	DBName      string
	Environment string
	CORSOrigins string
}

var AppConfig *Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		Port:        getEnv("PORT", "8080"),
		Host:        getEnv("HOST", "localhost"),
		DBDriver:    getEnv("DB_DRIVER", "sqlite"),
		DBName:      getEnv("DB_NAME", "notes.db"),
		Environment: getEnv("ENV", "development"),
		CORSOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetConfig() *Config {
	return AppConfig
}
