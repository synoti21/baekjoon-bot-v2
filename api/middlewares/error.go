package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors {
				if httpErr, ok := e.Err.(*errors.HTTPError); ok {
					ctx.String(httpErr.GetStatusCode(), httpErr.GetErrorMsg())
					return
				}
			}
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
