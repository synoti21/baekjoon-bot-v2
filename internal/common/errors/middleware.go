package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors {
				if httpErr, ok := e.Err.(*HTTPError); ok {
					ctx.String(httpErr.StatusCode, httpErr.Msg)
					return
				}
			}
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
