package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFeatures(c *gin.Context) {
	c.JSON(http.StatusOK, "feature: all features")
}
