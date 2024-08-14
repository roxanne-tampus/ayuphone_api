package main

import (
	"ayuphone_api/internal/controllers"
	"ayuphone_api/config"
	"ayuphone_api/internal/db"

	dbClient "ayuphone_api/pkg/db_client"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	newDbClient := dbClient.NewDBClient()
	defer newDbClient.DB.Close()

	dbService := db.NewDBService(newDbClient)

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	apiController := controllers.ApiController{
		DbClient:  newDbClient,
		DbService: dbService,
	}

	// Enable CORS
	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{env.Variables.FrontendURL, "http://localhost"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"} // Include the Authorization header
	router.Use(cors.New(config))

	// User Config
	router.POST("/api/apruv/get_user_info", apiController.UserInfo)
	router.POST("/api/apruv/login_for_testing", apiController.LoginForTesting)

	err := router.Run(":3001")
	if err != nil {
		return
	}

}
