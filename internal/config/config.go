package config

import (
	"os"
)

type Config struct {
	Port         string
	DBDSN        string
	JWTSecret    string
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPass     string
	FromEmail    string
	WebClientURL string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBDSN:        getEnv("DB_DSN", ""),
		JWTSecret:    getEnv("JWT_SECRET", "secret"),
		SMTPHost:     getEnv("SMTP_HOST", ""),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPass:     getEnv("SMTP_PASS", ""),
		FromEmail:    getEnv("FROM_EMAIL", "noreply@example.com"),
		WebClientURL: getEnv("WEB_CLIENT_URL", "http://localhost:5173"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
