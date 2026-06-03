package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func Events(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.GET("/", func(c *gin.Context) {
		data, e := st.ListEvents("")
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id", func(c *gin.Context) {
		data, e := st.GetEvent(c.Param("id"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.POST("/", guard, func(c *gin.Context) {
		var body store.EventInput
		if !api.Bind(c, &body) {
			return
		}
		data, e := st.CreateEvent(uid(c), body)
		if !err(c, e) {
			api.Created(c, data)
		}
	})
	r.PUT("/:id", guard, func(c *gin.Context) {
		var body store.EventInput
		if api.Bind(c, &body) {
			data, e := st.UpdateEvent(c.Param("id"), body)
			if !err(c, e) {
				api.OK(c, data)
			}
		}
	})
	r.DELETE("/:id", guard, func(c *gin.Context) {
		if !err(c, st.DeleteEvent(c.Param("id"))) {
			api.NoContent(c)
		}
	})
	seatRoutes(r, st, guard)
}

func seatRoutes(r *gin.RouterGroup, st store.Store, guard gin.HandlerFunc) {
	r.POST("/:id/seats/batch", guard, func(c *gin.Context) {
		var body store.AreaBatch
		if !api.Bind(c, &body) {
			return
		}
		data, e := st.CreateAreas(c.Param("id"), body)
		if !err(c, e) {
			api.Created(c, data)
		}
	})
	r.GET("/:id/seats", func(c *gin.Context) {
		data, e := st.ListSeats(c.Param("id"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id/seats/:seatId", func(c *gin.Context) {
		data, e := st.GetSeat(c.Param("id"), c.Param("seatId"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id/seats-status", guard, func(c *gin.Context) {
		data, e := st.SeatStatus(c.Param("id"), c.Query("eventDateId"), "")
		if !err(c, e) {
			api.OK(c, data)
		}
	})
	r.GET("/:id/seats-status/:seatId", guard, func(c *gin.Context) {
		data, e := st.SeatStatus(c.Param("id"), c.Query("eventDateId"), c.Param("seatId"))
		if !err(c, e) {
			api.OK(c, data)
		}
	})
}
