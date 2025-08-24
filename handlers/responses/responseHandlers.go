package response_handlers

import (
	"github.com/gin-gonic/gin"
)

func SuccessMessage(c *gin.Context, status int, msg any) {
	c.JSON(status, msg)
}
func FailMessage(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"error": msg,
	})
}
