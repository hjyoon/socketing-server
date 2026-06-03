package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
)

func uid(c *gin.Context) string {
	v, _ := c.Get("userId")
	id, _ := v.(string)
	return id
}

func err(c *gin.Context, e api.Error) bool {
	if e.Code == 0 {
		return false
	}
	api.Fail(c, e)
	return true
}
