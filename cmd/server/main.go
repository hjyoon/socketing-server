package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/hjyoon/socketing-server/internal/app"
	"github.com/hjyoon/socketing-server/internal/store"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		healthcheck()
		return
	}
	cfg := app.LoadConfig()
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	st := store.NewPostgres(db, cfg.JWTSecret)
	rc := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	defer rc.Close()
	st.SetRedis(rc)
	if err := st.EnsureSchema(); err != nil {
		log.Fatal(err)
	}
	if err := app.NewRouter(cfg, st).Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func healthcheck() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	c := http.Client{Timeout: 2 * time.Second}
	resp, err := c.Get("http://127.0.0.1:" + port + "/api/health")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusInternalServerError {
		os.Exit(1)
	}
}
