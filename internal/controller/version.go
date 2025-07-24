package controller

import "github.com/gin-gonic/gin"

const version = "1.0.0"

func Venson(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": version,
	})
}
