package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.Engine) {
	route.POST("/events", createEvent)
	route.POST("/update_event", updateEvent)
}
