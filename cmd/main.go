package main

import (
    "ayuphone_api/config"
    "ayuphone_api/internal/controllers"
    "ayuphone_api/internal/db"
    "ayuphone_api/internal/routes"
    "ayuphone_api/internal/services"
    "log"
    "os"

    "github.com/gin-contrib/cors"
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

    // Configure CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://10.0.2.2:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    routes.SetupRoutes(router, apiController)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    if err := router.Run("0.0.0.0:" + port); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}