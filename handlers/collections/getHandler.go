package collections_handlers

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	collections := []schemas.Collection{}
	if err := db.Find(&collections).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": collections,
	}
	response_handlers.Ok(c, msg)
}
