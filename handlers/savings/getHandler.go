package savings_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "WIP",
	})
}
