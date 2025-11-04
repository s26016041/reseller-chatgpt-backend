package ginrouter

import (
	"reseller-chatgpt-backend/internal/controller"
	"reseller-chatgpt-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, c *controller.Controller) {
	router.Use(middleware.StoreHeadersMiddleware())

	router.POST("/login", c.Login)

	router.GET("/version", c.Version)

	router.POST("/ask",
		middleware.VerifyJWT(),
		c.Ask)
}
