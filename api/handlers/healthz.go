package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *SlashCommandHandler) healthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Service Healthy")
	}
}
