package ginrouter

import (
	"reseller-chatgpt-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, c *controller.Controller) {
	router.GET("/version", c.Version)

	router.POST("/ask", c.Ask)
}
