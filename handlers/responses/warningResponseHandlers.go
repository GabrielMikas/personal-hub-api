package response_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context, msg any) {
	c.JSON(http.StatusNotFound, msg)
}

func BadRequest(c *gin.Context, msg any) {
	c.JSON(http.StatusBadRequest, msg)
}
