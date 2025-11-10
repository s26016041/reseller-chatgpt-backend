package main

import (
	"os"
	"reseller-chatgpt-backend/internal/controller"
	"reseller-chatgpt-backend/internal/env"
	"reseller-chatgpt-backend/internal/ginrouter"

	"github.com/gin-gonic/gin"
)

func main() {
	env.SetupEnv()

	router := ginrouter.Setting(gin.Default())

	ginrouter.SetupRoutes(router, controller.NewController())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
