package main

import (
	"net/http"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run("localhost:8080")
}

func getEvents(context *gin.Context)  {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context)  {
	var event models.Event
	if err := context.BindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	event.Save()
	context.JSON(http.StatusCreated, event)
}