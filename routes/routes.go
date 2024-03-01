package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routing for the server with the given gin Engine.
// It defines routes for event operations such as retrieving, creating, updating,
// and deleting events, as well as a signup route.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
