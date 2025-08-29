package collections_handlers

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetById(c *gin.Context) {
	id := c.Param("id")

	collection := schemas.Collection{}
	if err := db.Where("collection_id = ?", id).First(&collection).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	msg := gin.H{
		"data": collection,
	}
	response_handlers.SuccessMessage(c, http.StatusOK, msg)
}
