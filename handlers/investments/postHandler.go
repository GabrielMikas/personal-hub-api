package investments_handler

import (
	"fmt"
	"net/http"

	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostHandler(c *gin.Context) {
	req := Investment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	investment := schemas.Investment{
		InvestmentId: uuid.New().String(),
		Name:         req.Name,
		Type:         req.Type,
		Wallet:       req.Wallet,
		Amount:       req.Amount,
		Rate:         req.Rate,
		BoughtAt:     req.BoughtAt,
		ExpiresAt:    req.ExpiresAt,
		IsActive:     req.IsActive,
	}
	if err := db.Create(&investment).Error; err != nil {
		fmt.Println("Error creating Investment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"investment": req,
	})
}
