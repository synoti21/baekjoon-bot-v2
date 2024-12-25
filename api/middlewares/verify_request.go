package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
)

// Verify the HTTP request with signature header (or token header), using the adapter function
func VerifyRequestMiddleware(adapter adapters.Interface, signature string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if err := adapter.VerifyRequest(ctx.Request, signature); err != nil {
			ctx.AbortWithError(err.GetStatusCode(), err)
		}
	}
}
