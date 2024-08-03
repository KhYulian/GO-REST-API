package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"rest-api/utils"
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

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized"})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

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

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse reuqest data. Error message is: %s", err)})
		return
	}

	_, err = models.GetOneEvent(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not fetch the event. Error message is: %s", err)})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid body. Error message: %s", err)})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not update the event. Error message: %s", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse reuqest data. Error message is: %s", err)})
		return
	}

	event, err := models.GetOneEvent(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not fetch the event. Error message is: %s", err)})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not delete the event. Error message is: %s", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}
