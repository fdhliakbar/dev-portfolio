package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Environment string
	Port        int
	ClientURL   string
	DatabaseURL string
	JWTSecret   string
}

// New creates a new Config instance from environment variables
func New() *Config {
	port, err := strconv.Atoi(getEnvOrDefault("APP_PORT", "8080"))
	if err != nil {
		port = 8080
	}

	return &Config{
		Environment: getEnvOrDefault("APP_ENV", "development"),
		Port:        port,
		ClientURL:   getEnvOrDefault("CLIENT_URL", "http://localhost:3000"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   getEnvOrDefault("JWT_SECRET", "dev-secret-key"),
	}
}

// Helper function to get environment variable with a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}