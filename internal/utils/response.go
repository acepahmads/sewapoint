package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cara pakai
// utils.Success(c, data)
// utils.Error(c, 400, "invalid request")

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// Error response
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Success: false,
		Error: gin.H{
			"message": message,
		},
	})
}
