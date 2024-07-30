package main

import (
	"fmt"
	"net/http"
	"os"

	"rest-api/db"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("CGO_ENABLED", "1")
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET POST PUTE PATCH DELETE etc
	server.POST("/events", createEvent)

	server.Run(":8080") // PORT 8080 on localhost
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events) // will be automatically transformed to the JSON by the GIN package
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

	event.Save()
	context.JSON(http.StatusOK, gin.H{"message": "Event created"})
}
