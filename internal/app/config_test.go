package app

import (
	"os"
	"testing"
)

func TestLoadConfigDefaultsAndEnv(t *testing.T) {
	t.Setenv("PORT", "9000")
	t.Setenv("JWT_SECRET", "secret")
	t.Setenv("DB_URL", "postgres://x")
	t.Setenv("CORS_ALLOWED_ORIGINS", "http://a, http://b")
	cfg := LoadConfig()
	if cfg.Port != "9000" || cfg.JWTSecret != "secret" || cfg.DBURL != "postgres://x" {
		t.Fatalf("env config mismatch")
	}
	if len(cfg.CORSOrigins) != 2 || cfg.CORSOrigins[1] != "http://b" {
		t.Fatalf("cors origins mismatch")
	}
	os.Unsetenv("DB_URL")
	t.Setenv("DB_HOST", "h")
	if buildDBURL() == "" || env("MISSING", "fallback") != "fallback" ||
		splitList(" , ")[0] != "*" {
		t.Fatalf("default config failed")
	}
}
