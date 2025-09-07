package cards_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	cards := []schemas.Card{}
	if err := db.Find(&cards).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": cards,
	}
	response_handlers.Ok(c, msg)
}
