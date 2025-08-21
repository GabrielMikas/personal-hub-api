package investments_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	var req Investment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"investment": req,
	})
}
