package handlers

import (
	"github.com/GabrielMikas/personal-hub-api/config"
	cards_handler "github.com/GabrielMikas/personal-hub-api/handlers/cards"
	collections_handlers "github.com/GabrielMikas/personal-hub-api/handlers/collections"
	investments_handler "github.com/GabrielMikas/personal-hub-api/handlers/investments"
	savings_handler "github.com/GabrielMikas/personal-hub-api/handlers/savings"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandlers() {
	db = config.GetDB()
	investments_handler.InitInvestmentsHandler(db)
	savings_handler.InitSavingsHandler(db)
	cards_handler.InitCardsHandler(db)
	collections_handlers.InitCollectionsHandler(db)
}
