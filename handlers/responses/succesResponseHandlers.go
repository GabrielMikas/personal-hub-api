package response_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, msg any) {
	c.JSON(http.StatusOK, msg)
}
func Created(c *gin.Context, msg any) {
	c.JSON(http.StatusCreated, msg)
}
func Accepted(c *gin.Context, msg any) {
	c.JSON(http.StatusAccepted, msg)
}
func NoContent(c *gin.Context, msg any) {
	c.JSON(http.StatusNoContent, msg)
}
