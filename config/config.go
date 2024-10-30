package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type appConfig struct {
	Port    string
	Mode    string
	BaseURL string
}

type dbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type jwtConfig struct {
	SecretKey                   string
	TokenExpiration             time.Duration
	ResetTokenExpiration        time.Duration
	VerificationTokenExpiration time.Duration
}

type smtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

var (
	App  *appConfig
	DB   *dbConfig
	JWT  *jwtConfig
	SMTP *smtpConfig
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App = &appConfig{
		Port:    getEnv("APP_PORT"),
		Mode:    getEnv("APP_MODE"),
		BaseURL: getEnv("APP_BASE_URL"),
	}

	DB = &dbConfig{
		Host:     getEnv("DB_HOST"),
		Port:     getEnv("DB_PORT"),
		Username: getEnv("DB_USERNAME"),
		Password: getEnv("DB_PASSWORD"),
		Name:     getEnv("DB_NAME"),
	}

	JWT = &jwtConfig{
		SecretKey:                   getEnv("JWT_SECRET_KEY"),
		TokenExpiration:             getEnvAsDuration("JWT_TOKEN_EXPIRATION", "15m"),
		ResetTokenExpiration:        getEnvAsDuration("JWT_RESET_TOKEN_EXPIRATION", "60m"),
		VerificationTokenExpiration: getEnvAsDuration("JWT_VERIFICATION_TOKEN_EXPIRATION", "1440m"),
	}

	SMTP = &smtpConfig{
		Host:     getEnv("SMTP_HOST"),
		Port:     getEnv("SMTP_PORT"),
		Username: getEnv("SMTP_USERNAME"),
		Password: getEnv("SMTP_PASSWORD"),
		From:     getEnv("SMTP_FROM"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Environment variable %s is required", key)
	return ""
}

/* func getEnvAsInt(key string) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	log.Fatalf("Environment variable %s must be an integer", key)
	return 0
} */

func getEnvAsDuration(key, defaultVal string) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	duration, _ := time.ParseDuration(defaultVal)
	return duration
}
