package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/auth"
	"github.com/hjyoon/socketing-server/internal/routes"
	"github.com/hjyoon/socketing-server/internal/store"
)

func NewRouter(cfg Config, st store.Store) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	origins := cfg.CORSOrigins
	if len(origins) == 0 {
		origins = []string{"*"}
	}
	r.Use(gin.Logger(), gin.Recovery(), cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "Accept", "Authorization",
		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := r.Group("/api")
	v1.GET("/health", func(c *gin.Context) {
		if err := st.Health(); err != nil {
			api.Fail(c, api.ErrInternal)
			return
		}
		api.OK(c, map[string]string{"status": "ok"})
	})
	guard := auth.Middleware(cfg.JWTSecret)
	routes.Auth(v1.Group("/auth"), st, cfg.JWTSecret)
	routes.Users(v1.Group("/users"), st, guard)
	routes.Events(v1.Group("/events"), st, guard)
	routes.Orders(v1.Group("/orders"), st, guard)
	routes.Payments(v1.Group("/payments"), st, guard)
	routes.Managers(v1.Group("/managers/events"), st, guard)
	return r
}
