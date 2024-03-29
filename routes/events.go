package routes

import (
	"net/http"
	"strconv"

	"github.com/EtienneRQT/Go_REST_API/models"
	"github.com/gin-gonic/gin"
)

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
		return
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
		return
	}
	context.JSON(http.StatusOK, events)
}

// createEvent handles POST requests to create a new event.
// It first verifies the JWT authorization token.
// It then binds the event data from the request body.
// The event is saved to the database, returning 201 Created if successful,
// or an error response if there is an error binding or saving the event.
func createEvent(context *gin.Context) {

	var event models.Event
	if err := context.BindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userID := context.GetInt64("userId")
	event.UserID = userID
	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, event)
}

// updateEvent handles PUT requests to update an existing event.
// It parses the event ID from the URL parameter, binds the request body to an Event struct,
// validates the event exists, updates it in the database, and returns the updated event.
// It returns 404 if the event is not found, 400 for binding or validation errors,
// or 500 for any database errors.
func updateEvent(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	userID, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	}

	if event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	var updatedEvent models.Event
	if err := context.BindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, updatedEvent)
}

// deleteEvent handles DELETE requests to delete an event by ID.
// It parses the event ID from the URL parameter, looks up the event,
// deletes it from the database, and returns a 200 OK response.
// Returns 404 if the event is not found, or 500 for any database errors.
func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.Delete(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
