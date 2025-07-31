package middleware

import (
	"github.com/gin-gonic/gin"
)

func StoreHeadersMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for key, values := range ctx.Request.Header {
			if len(values) > 0 {
				ctx.Set(key, values[0])
			}
		}
		ctx.Next()
	}
}
