package collections_handlers

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PatchHandler(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}

	if err := c.ShouldBindJSON(&updates); err != nil {
		response_handlers.FailMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Model(&schemas.Collection{}).Where("collection_id = ?", id).Updates(&updates).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg := gin.H{
		"message":      "Collection updated successfully",
		"collectionId": id,
	}
	response_handlers.SuccessMessage(c, http.StatusOK, msg)
}
