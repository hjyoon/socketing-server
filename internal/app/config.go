package app

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port        string
	JWTSecret   string
	DBURL       string
	CORSOrigins []string
}

func LoadConfig() Config {
	port := env("PORT", "5000")
	return Config{
		Port:      port,
		JWTSecret: env("JWT_SECRET", "my-jwt-secret"),
		DBURL:     env("DB_URL", buildDBURL()),
		CORSOrigins: splitList(
			env("CORS_ALLOWED_ORIGINS", "*"),
		),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func buildDBURL() string {
	host := env("DB_HOST", "localhost")
	port := env("DB_PORT", "5432")
	user := env("DB_USERNAME", "postgres")
	pass := env("DB_PASSWORD", "password")
	name := env("DB_NAME", "socketing")
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name,
	)
}

func splitList(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if item := strings.TrimSpace(part); item != "" {
			out = append(out, item)
		}
	}
	if len(out) == 0 {
		return []string{"*"}
	}
	return out
}
