package investments_handler

import (
	response_handlers "github.com/GabrielMikas/personal-hub-api/handlers/responses"
	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostHandler(c *gin.Context) {
	req := Investment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response_handlers.BadRequest(c, err.Error())
		return
	}
	investment := schemas.Investment{
		InvestmentId: uuid.New().String(),
		Name:         req.Name,
		Description:  req.Description,
		Type:         req.Type,
		Wallet:       req.Wallet,
		Amount:       req.Amount,
		Rate:         req.Rate,
		Liquidity:    req.Liquidity,
		BoughtAt:     req.BoughtAt,
		ExpiresAt:    req.ExpiresAt,
		IsActive:     req.IsActive,
	}
	if err := db.Create(&investment).Error; err != nil {
		response_handlers.InternalServerError(c, err.Error())
		return
	}
	msg := gin.H{
		"message":      "Investment created successfully",
		"investmentId": investment.InvestmentId,
	}
	response_handlers.Created(c, msg)

}
