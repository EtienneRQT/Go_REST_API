package main

import (
	"example.com/api/db"
	"example.com/api/routes"
	"github.com/gin-gonic/gin"
)

// main is the entry point for the application. It initializes the
// database, sets up the HTTP routes, and starts the server.
func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run("localhost:8080")
}
