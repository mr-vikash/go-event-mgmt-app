package routes

import (
	"net/http"
	"strconv"

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

func updateEvent(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event id",
		})
		return
	}

	var req models.Event
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	event, err := repository.UpdateEvent(id, req.Name, req.Location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

func deleteEvent(cxt *gin.Context) {

	idParam := cxt.Param("id")

	id, err := strconv.Atoi(idParam)

	result, err := repository.DeleteEvent(id)

	if err != nil {
		cxt.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": result,
		})
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": result,
	})
}

func getEvents(context *gin.Context) {
	events, err := repository.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Events loaded successfully",
		"events":  events,
	})
}
