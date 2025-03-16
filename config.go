package io_base_service

import (
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold environment variables
type Config struct {
	Port      string
	LogLevel  string
	LogFolder string
	EnvType   string
	JWTSecret string
}

// LoadConfig initializes and loads environment variables
func LoadConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	envFile := ".env"
	switch env {
	case "production":
		envFile = ".env.production"
	case "staging":
		envFile = ".env.staging"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		Warn("No .env file found, using system environment variables")
	}

	logFolder := getEnv("LOG_FOLDER", "logs")
	logLevel := getEnv("LOG_LEVEL", "info")

	// Initialize Logger
	InitLogger(logFolder, logLevel)

	Info("Application environment loaded: %v", map[string]interface{}{"environment": env})

	return &Config{
		Port:      getEnv("PORT", "3030"),
		LogLevel:  logLevel,
		LogFolder: logFolder,
		EnvType:   env,
		JWTSecret: getEnv("JWT_SECRET", "default-secret-key"),
	}
}

// Helper function to read environment variables with a default fallback
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
