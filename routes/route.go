package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.Engine) {
	route.POST("/events", createEvent)
	route.PUT("/events/:id", updateEvent)
	route.DELETE("/events/:id", deleteEvent)
	route.GET("/events", getEvents)
}
