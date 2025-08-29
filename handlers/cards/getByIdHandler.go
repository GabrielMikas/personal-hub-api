package cards_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetById(c *gin.Context) {
	id := c.Param("id")

	card := schemas.Card{}
	if err := db.Where("card_id = ?", id).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	msg := gin.H{
		"data": card,
	}
	response_handlers.SuccessMessage(c, http.StatusOK, msg)
}
