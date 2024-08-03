package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents) // GET POST PUTE PATCH DELETE etc
	server.GET("/events/:eventID", getOneEvent)

	// AUTH
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Private endpoints
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventID", updateEvent)
	authenticated.DELETE("/events/:eventID", deleteEvent)

}
