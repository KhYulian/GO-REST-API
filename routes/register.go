package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("user_id")
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id param"})
		return
	}

	event, err := models.GetOneEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("user_id")
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id param"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}
