package controller

import (
	"reseller-chatgpt-backend/internal/env"

	"github.com/gin-gonic/gin"
)

func Ask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"apikey": env.GetOpenAIAPIKey(),
	})
}
