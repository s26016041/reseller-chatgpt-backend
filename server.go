package main

import (
	"reseller-chatgpt-backend/internal/controller"
	"reseller-chatgpt-backend/internal/env"
	"reseller-chatgpt-backend/internal/ginrouter"

	"github.com/gin-gonic/gin"
)

func main() {
	env.SetupEnv()

	router := gin.Default()

	ginrouter.SetupRoutes(router, controller.NewController())

	router.Run(":8080")

}
