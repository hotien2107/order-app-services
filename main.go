package main

import (
	"net/http"
	"order-app/web-service-gin/database"
	"order-app/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func getFoodMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, foodMenu)
}

func main() {
	router := gin.Default()
	database.Connect()
	router.Use(CORSMiddleware())

	router.GET("/food-menu", getFoodMenu)
	routes.FoodRoute(router)

	router.Run()
}
