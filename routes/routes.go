package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public endpoints
	// Events
	server.GET("/events", getEvents) // GET POST PUTE PATCH DELETE etc
	server.GET("/events/:eventID", getOneEvent)

	// AUTH
	server.POST("/signup", signup)
	server.POST("/login", login)
	// END of Public endpoints

	// Private endpoints
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Events
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventID", updateEvent)
	authenticated.DELETE("/events/:eventID", deleteEvent)

	// Registrations
	authenticated.POST("/events/:eventID/register", registerForEvent)
	authenticated.DELETE("/events/:eventID/register", cancelRegistration)
}
