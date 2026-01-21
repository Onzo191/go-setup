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
	Env      string
	Port     string
	DB       DBConfig
	Security SecurityConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type SecurityConfig struct {
	JwtSecret          string
	RefreshTokenSecret string
	PasswordHashSecret string
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
		Security: SecurityConfig{
			JwtSecret:          GetEnv("JWT_SECRET", "default_jwt_secret"),
			RefreshTokenSecret: GetEnv("REFRESH_TOKEN_SECRET", "default_refresh_token_secret"),
			PasswordHashSecret: GetEnv("PASSWORD_HASH_SECRET", "default_password_hash_secret"),
		},
	}

	return cfg
}
