package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
)

func Middleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		if token == header || token == "" {
			api.Fail(c, api.ErrUnauthorized)
			c.Abort()
			return
		}
		claims, err := Verify(token, secret)
		if err != nil {
			api.Fail(c, api.ErrUnauthorized)
			c.Abort()
			return
		}
		c.Set("userId", claims.Sub)
		c.Next()
	}
}
