package main

import (
	"ayuphone_api/config"
	"ayuphone_api/internal/controllers"
	"ayuphone_api/internal/db"
	"ayuphone_api/internal/routes"
	"ayuphone_api/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	dbClient, err := db.NewSQLiteDBClient()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	dbService := services.NewDBService(&dbClient)

	apiController := controllers.ApiController{
		DbClient:  &dbClient,
		DbService: dbService,
	}

	router := gin.Default()
	routes.SetupRoutes(router, apiController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
