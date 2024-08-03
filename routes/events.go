package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An error occured while fetching events. Error message: %s", err)})
		return
	}

	context.JSON(http.StatusOK, events) // will be automatically transformed to the JSON by the GIN package
}

func getOneEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid parameter for eventID. Error message: %s", err)})
		return
	}

	event, err := models.GetOneEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not fetch the event. Error message: %s", err)})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	// body := context.Request.Body
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse reuqest data. Error message is: %s", err)})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not create an event. Error message is: %s", err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event created"})
}
