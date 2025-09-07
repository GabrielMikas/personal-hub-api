package investments_handler

import (
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
			response_handlers.NotFound(c, err)
			return
		}
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"data": investment,
	}
	response_handlers.Ok(c, msg)
}
