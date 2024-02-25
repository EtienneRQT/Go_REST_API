package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
server.GET("/events", getEvents)	

	server.Run("localhost:8080")
}

func getEvents(context *gin.Context)  {
	context.JSON(200, gin.H{
		"events": []string{"Meeting", "Lunch", "Party"},
	})
}