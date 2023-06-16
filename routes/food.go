package routes

import (
	"order-app/web-service-gin/controller"

	"github.com/gin-gonic/gin"
)

func FoodRoute(router *gin.Engine) {
	router.GET("/api/foodGroup", controller.GetFoodGroup)
	router.POST("/api/foodGroup", controller.InsertFoodGrooup)
	router.PUT("/api/foodGroup", controller.UpdateFoodGroup)
	router.DELETE("/api/foodGroup", controller.DeleteFoodGroup)
}
