package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, statusCode int, message string) {
	context.Header("Content-type", "application/json")
	context.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
	})
}

func sendSuccess(context *gin.Context, operation string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", operation),
		"data":    data,
	})
}
