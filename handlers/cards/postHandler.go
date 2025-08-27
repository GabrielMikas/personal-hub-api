package cards_handler

import (
	"net/http"

	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PostHandler(c *gin.Context) {
	req := Card{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response_handlers.FailMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	collection := schemas.Collection{}
	if err := db.Where("collection_id = ?", req.CollectionId).First(&collection).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response_handlers.FailMessage(c, http.StatusNotFound, err.Error())
			return
		}
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	card := schemas.Card{
		CardId:       uuid.New().String(),
		Name:         req.Name,
		Description:  req.Description,
		CollectionId: req.CollectionId,
		Type:         req.Type,
		CardCode:     req.CardCode,
		Amount:       req.Amount,
		Rarity:       req.Rarity,
		BoughtAt:     req.BoughtAt,
		IsSleeved:    req.IsSleeved,
	}

	if err := db.Create(&card).Error; err != nil {
		response_handlers.FailMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	msg := gin.H{
		"message": "Card created successfully",
		"cardId":  card.CardId,
	}
	response_handlers.SuccessMessage(c, http.StatusCreated, msg)

}
