package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
)

// ErrorHandlingMiddleware is middleware to get the consistent error response structure
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors {
				if httpErr, ok := e.Err.(*errors.BaseError); ok {
					ctx.JSON(httpErr.GetStatusCode(), gin.H{"error": httpErr.GetErrorMsg()})
					return
				}
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ctx.Errors.Errors()})
		}
	}
}
