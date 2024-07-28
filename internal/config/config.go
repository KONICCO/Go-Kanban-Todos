package config

import (
	"os"
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
	db := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", db),
		},
		Server: ServerConfig{
			Address: ":" + port,
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
