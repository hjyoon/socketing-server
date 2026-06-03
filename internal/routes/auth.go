package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Auth(r *gin.RouterGroup, st store.Store, secret string) {
	r.POST("/register", func(c *gin.Context) {
		var body store.Register
		if !api.Bind(c, &body) {
			return
		}
		data, e := st.CreateUser(body)
		if err(c, e) {
			return
		}
		api.Created(c, data)
	})
	r.POST("/login", func(c *gin.Context) {
		var body store.Login
		if !api.Bind(c, &body) {
			return
		}
		data, e := st.Login(body, secret)
		if err(c, e) {
			return
		}
		api.OK(c, data)
	})
}
