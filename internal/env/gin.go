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
