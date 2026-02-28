package routes

import (
	"net/http"

	"go-event-mgmt-app/models"
	"go-event-mgmt-app/repository"

	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
	}

	eventId := repository.InsertEvent(event.Name, event.Location)

	context.JSON(http.StatusOK, gin.H{
		"message":  "Event created successfully",
		"event_id": eventId,
	})
}

func updateEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusSeeOther, gin.H{
			"error": "Something went wrong",
		})
	}

	event, err = repository.UpdateEvent(event.ID, event.Name, event.Location)

	if err != nil {
		context.JSON(http.StatusSeeOther, gin.H{
			"error": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})

}
