package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthz is a health check function
func (h *SlashCommandHandler) healthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Service Healthy")
	}
}
