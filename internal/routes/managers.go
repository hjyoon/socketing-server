package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Managers(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.GET("/", guard, func(c *gin.Context) {
		data, e := st.ListEvents(uid(c))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:eventId/reservation-status", guard, func(c *gin.Context) {
		data, e := st.ManagerEvent(uid(c), c.Param("eventId"), c.Query("eventDateId"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
}
