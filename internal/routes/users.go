package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Users(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.GET("/email/:email", func(c *gin.Context) {
		data, e := st.GetUserByEmail(c.Param("email"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id/points", func(c *gin.Context) {
		data, e := st.GetPoints(c.Param("id"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id", func(c *gin.Context) {
		data, e := st.GetUser(c.Param("id"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.PATCH("/:id/nickname", guard, func(c *gin.Context) {
		var b struct {
			Nickname string `json:"nickname"`
		}
		if !api.Bind(c, &b) {
			return
		}
		data, e := st.UpdateNickname(c.Param("id"), b.Nickname)
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.PATCH("/:id/password", guard, func(c *gin.Context) {
		var b struct {
			Password string `json:"password"`
		}
		if api.Bind(c, &b) && !err(c, st.UpdatePassword(c.Param("id"), b.Password)) {
			api.OK(c, nil)
		}
	})
	r.DELETE("/:id", guard, func(c *gin.Context) {
		if !err(c, st.DeleteUser(c.Param("id"))) {
			api.NoContent(c)
		}
	})
}
