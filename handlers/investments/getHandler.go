package investments_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	investments := []schemas.Investment{}
	if err := db.Find(&investments).Error; err != nil {
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
	}
	msg := gin.H{
		"data": investments,
	}
	response_handlers.SuccessMessage(c, http.StatusOK, msg)
}
