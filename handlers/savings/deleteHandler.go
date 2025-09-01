package savings_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("saving_id = ?", id).Delete(&schemas.Saving{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.NotFound(c, err.Error())
			return
		}
		response_handlers.InternalServerError(c, err.Error())
	}
	msg := gin.H{
		"message":  "Saving deleted succesfully",
		"savingId": id,
	}
	response_handlers.Accepted(c, msg)
}
