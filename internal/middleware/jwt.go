package middleware

import (
	"net/http"
	"reseller-chatgpt-backend/internal/constant"
	"reseller-chatgpt-backend/internal/env"
	"reseller-chatgpt-backend/internal/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyJWT() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		authorization := strings.TrimPrefix(env.GetAuthorization(ctx), "Bearer ")
		jwtClaims, err := utils.ParseJWT(authorization)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Set(constant.JWTUsername, jwtClaims.Username)
		ctx.Set(constant.JWTPassword, jwtClaims.Password)

		ctx.Next()
	}
}
