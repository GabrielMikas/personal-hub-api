package savings_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetById(c *gin.Context) {
	id := c.Param("id")

	saving := schemas.Saving{}
	if err := db.Where("saving_id = ?", id).First(&saving).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.NotFound(c, err.Error())
			return
		}
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": saving,
	}
	response_handlers.Ok(c, msg)
}
