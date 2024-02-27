package routes

import (
	"github.com/gin-gonic/gin"
)


// RegisterRoutes registers the routes for the Gin server.
// It registers the following routes:
// - GET /events: Handles getting all events
// - GET /events/:id: Handles getting a single event by ID
// - POST /events: Handles creating a new event
// - PUT /events/:id: Handles updating an existing event by ID
// - DELETE /events/:id: Handles deleting an event by ID
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
