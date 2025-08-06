package env

import (
	"reseller-chatgpt-backend/internal/constant"

	"github.com/gin-gonic/gin"
)

func GetAuthorization(ctx *gin.Context) string {
	authorization, exists := ctx.Get(constant.Authorization)
	if !exists {
		return ""
	}

	return authorization.(string)
}

func GetJWTUsername(ctx *gin.Context) string {
	username, exists := ctx.Get(constant.JWTUsername)
	if !exists {
		return ""
	}

	return username.(string)
}

func GetJWTPassword(ctx *gin.Context) string {
	password, exists := ctx.Get(constant.JWTPassword)
	if !exists {
		return ""
	}

	return password.(string)
}
