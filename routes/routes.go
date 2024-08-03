package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents) // GET POST PUTE PATCH DELETE etc
	server.GET("/events/:eventID", getOneEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:eventID", updateEvent)
	server.DELETE("/events/:eventID", deleteEvent)
}
