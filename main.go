package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET POST PUTE PATCH DELETE etc

	server.Run(":8080") // PORT 8080 on localhost
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Heello!"})
}
