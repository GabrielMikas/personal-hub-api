package response_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, msg any) {
	c.JSON(http.StatusInternalServerError, msg)
}
