package savings_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PatchHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("saving_id = ?", id).Update("is_active", gorm.Expr("NOT is_active")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.NotFound(c, err)
			return
		}
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"message":   "Toggled saving active status successfully",
		"saving_id": id,
	}
	response_handlers.Ok(c, msg)
}
