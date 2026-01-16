package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	envPath := ".env"
	if path != "" {
		envPath = path
	}

	err := godotenv.Load(envPath)
	if err != nil {
		println("env:cannot load .env file: " + err.Error())
	}
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

type Config struct {
	Env  string
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadEnvConfig(path string) *Config {
	LoadEnv(path)

	cfg := &Config{
		Env:  GetEnv("ENV", "development"),
		Port: GetEnv("PORT", "8080"),
		DB: DBConfig{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "password"),
			DBName:   GetEnv("DB_NAME", "mydb"),
		},
	}

	return cfg
}
