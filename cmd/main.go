package main

import (
	"log"
	"os"

	"ayuphone_api/config"
	db "ayuphone_api/internal/db"
	"ayuphone_api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
