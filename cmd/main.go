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
	router.Use(CORSMiddleware())
	routes.SetupRoutes(router, apiController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
