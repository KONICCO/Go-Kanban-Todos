package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	URL string
}

type ServerConfig struct {
	Address string
}

func New() *Config {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	db := os.Getenv("DATABASE_URL")
	print(db)
	return &Config{

		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", db),
		},
		Server: ServerConfig{
			Address: getEnv("PORT", ":8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
