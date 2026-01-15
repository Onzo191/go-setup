package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadEnv(path string) *Config {
	var envPath string = ".env"
	if path != "" {
		envPath = path
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		PORT: GetEnv("PORT", "8080"),
	}

	log.Printf("Configuration loaded: %+v\n", cfg)
	return cfg
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
