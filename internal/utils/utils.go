package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/logger"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger.LogError(message, "Error Service Resonce")
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

func NewSuccessResponce(ctx *gin.Context, body interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"body": body,
	})
}
