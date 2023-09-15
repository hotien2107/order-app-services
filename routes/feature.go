package routes

import (
	"order-app/web-service-gin/controller"

	"github.com/gin-gonic/gin"
)

func FeatureRoute(router *gin.Engine) {
	router.POST("/getAllFeatures", controller.GetAllFeatures)
}
