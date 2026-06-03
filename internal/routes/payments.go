package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Payments(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.POST("/", guard, func(c *gin.Context) {
		var body store.PaymentInput
		if !api.Bind(c, &body) {
			return
		}
		data, e := st.CreatePayment(uid(c), body)
		if !err(c, e) {
			api.Created(c, data)
		}
	})
	r.PATCH("/", guard, func(c *gin.Context) {
		api.Created(c, nil)
	})
}
