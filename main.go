package main

import (
	"net/http"
	"strconv"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

// main is the entry point for the application. It initializes the
// database, sets up the HTTP routes, and starts the server.
func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.Run("localhost:8080")
}

// getEvent retrieves an event by ID from the database.
// It parses the ID from the URL parameter, queries the database,
// and returns the event as JSON. If there is an error parsing the ID
// or querying the database, it returns an error response.
func getEvent(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, event)
}

// getEvents handles GET requests to retrieve all events.
// It calls the GetAllEvents model method to fetch events from the database.
// If there is an error, it returns a 500 status code and the error message.
// Otherwise, it returns a 200 status code and the list of events.
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, events)
}

// createEvent handles POST requests to create a new event.
// It binds the request body to an Event struct, validates the data,
// saves the event to the database, and returns the created event
// with 201 Created status code on success, or error responses on failure.
func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.BindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, event)
}
