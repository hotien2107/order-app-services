package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getFoodMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, foodMenu)
}

func main() {
	router := gin.Default()
	router.GET("/food-menu", getFoodMenu)

	router.Run("localhost:8080")
}
