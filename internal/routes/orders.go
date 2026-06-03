package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Orders(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.GET("/", guard, func(c *gin.Context) {
		data, e := st.ListOrders(uid(c), c.Query("eventId"))
		if !err(c, e) {
			api.Created(c, data)
		}
	})
	r.GET("/:orderId", guard, func(c *gin.Context) {
		data, e := st.GetOrder(c.Param("orderId"), uid(c))
		if !err(c, e) {
			api.Created(c, data)
		}
	})
	r.POST("/:orderId/cancel", guard, func(c *gin.Context) {
		if !err(c, st.CancelOrder(c.Param("orderId"), uid(c))) {
			api.Created(c, map[string]any{})
		}
	})
}
