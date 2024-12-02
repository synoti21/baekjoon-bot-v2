package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
)

func VerifyRequestMiddleware(adapter adapters.Interface, signature string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if err := adapter.VerifyRequest(ctx.Request, signature); err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.NewUnauthorizedError(err.Error()))
		}
	}
}
