package main

import (
	"order-app/web-service-gin/database"
	"order-app/web-service-gin/helpers"
	"order-app/web-service-gin/middleware"
	"order-app/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.GetEnvironmentVariable()

	router := gin.Default()
	database.Connect()
	router.Use(middleware.CORSMiddleware())
	routes.FoodRoute(router)
	routes.FeatureRoute(router)

	router.Run()
}
