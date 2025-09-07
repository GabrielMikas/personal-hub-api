package savings_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	savings := []schemas.Saving{}
	if err := db.Find(&savings).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": savings,
	}
	response_handlers.Ok(c, msg)
}
