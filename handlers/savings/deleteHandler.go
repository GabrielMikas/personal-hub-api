package savings_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("saving_id = ?", id).Delete(&schemas.Saving{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
	}
	msg := gin.H{
		"message":  "Saving deleted succesfully",
		"savingId": id,
	}
	response_handlers.SuccessMessage(c, http.StatusAccepted, msg)
}
