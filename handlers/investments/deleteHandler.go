package investments_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("investment_id = ?", id).Delete(&schemas.Investment{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.NotFound(c, err)
			return
		}
		response_handlers.InternalServerError(c, err.Error())
	}
	msg := gin.H{
		"message":      "Investment deleted succesfully",
		"investmentId": id,
	}
	response_handlers.Accepted(c, msg)
}
