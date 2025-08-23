package investments_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetById(c *gin.Context) {
	id := c.Param("id")

	investment := schemas.Investment{}
	if err := db.Where("investment_id = ?", id).First(&investment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	msg := gin.H{
		"data": investment,
	}
	response_handlers.SuccessMessage(c, http.StatusOK, msg)
}
