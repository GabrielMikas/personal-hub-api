package investments_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("investment_id = ?", id).Delete(&schemas.Investment{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
	}
	msg := gin.H{
		"message":      "Investment deleted succesfully",
		"investmentId": id,
	}
	response_handlers.SuccessMessage(c, http.StatusAccepted, msg)
}
