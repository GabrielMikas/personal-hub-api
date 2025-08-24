package handlers

import (
	"github.com/GabrielMikas/personal-hub-api/config"
	investments_handler "github.com/GabrielMikas/personal-hub-api/handlers/investments"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandlers() {
	db = config.GetDB()
	investments_handler.InitInvestmentsHandler(db)
}
