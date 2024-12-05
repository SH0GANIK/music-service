package handler

import (
	"github.com/gin-gonic/gin"
	"music-service/internal/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Log.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
