package investments_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	investments := []schemas.Investment{}
	if err := db.Find(&investments).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": investments,
	}
	response_handlers.Ok(c, msg)
}
