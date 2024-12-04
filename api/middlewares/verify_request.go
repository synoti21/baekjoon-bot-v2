package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
)

func VerifyRequestMiddleware(adapter adapters.Interface, signature string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if err := adapter.VerifyRequest(ctx.Request, signature); err != nil {
			ctx.AbortWithError(err.GetStatusCode(), err)
		}
	}
}
