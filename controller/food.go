package controller

import (
	"net/http"
	"order-app/web-service-gin/database"
	"order-app/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func GetFoodGroup(c *gin.Context) {
	foodGroup := []models.FoodGroup{}
	database.DB.Find(&foodGroup)
	c.JSON(http.StatusOK, &foodGroup)
}

func InsertFoodGrooup(c *gin.Context) {
	var foodGroup models.FoodGroup
	c.BindJSON(&foodGroup)
	result := database.DB.Create(&foodGroup)
	if result.Error != nil {
		c.String(http.StatusOK, result.Error.Error())
	}
	c.JSON(http.StatusOK, &foodGroup)
}

func UpdateFoodGroup(c *gin.Context) {
	var foodGroup models.FoodGroup
	c.BindJSON(&foodGroup)
	result := database.DB.Save(&foodGroup)
	if result.Error != nil {
		c.String(http.StatusOK, result.Error.Error())
	}
	c.JSON(http.StatusOK, &foodGroup)
}

func DeleteFoodGroup(c *gin.Context) {
	foodGroupID := c.Param("id")
	var foodGroup models.FoodGroup
	result := database.DB.Where("id = ?", foodGroupID).Delete(&foodGroup)
	if result.Error != nil {
		c.String(http.StatusOK, result.Error.Error())
	}
	c.String(http.StatusOK, "Delete food group success")
}
