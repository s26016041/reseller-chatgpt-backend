package ginrouter

import (
	"reseller-chatgpt-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/version", controller.Venson)

	router.GET("/ask", controller.Ask)

}
