package cards_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	var req Card
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"card": req,
	})
}
