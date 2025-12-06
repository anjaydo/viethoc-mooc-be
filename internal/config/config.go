package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	MuxTokenID  string
	MuxSecret   string
}

// LoadConfig reads .env file or environment variables
func LoadConfig() *Config {
	// Attempt to load .env file, but don't fail if it doesn't exist (prod envs)
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""), // Connection string for Supabase
		MuxTokenID:  getEnv("MUX_TOKEN_ID", ""),
		MuxSecret:   getEnv("MUX_TOKEN_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if fallback == "" && key != "PORT" {
		log.Printf("Warning: %s is not set", key)
	}
	return fallback
}
