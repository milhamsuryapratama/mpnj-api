package utils

import (
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, message string, data interface{}, status int) {
	c.JSON(status, gin.H{
		"message": message,
		"data": data,
	})
}