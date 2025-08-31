package response_handlers

import (
	"net/http"

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
func NotFound(c *gin.Context, msg any) {
	c.JSON(http.StatusNotFound, msg)
}
func InternalServerError(c *gin.Context, msg any) {
	c.JSON(http.StatusInternalServerError, msg)
}
func BadRequest(c *gin.Context, msg any) {
	c.JSON(http.StatusBadRequest, msg)
}
