package cards_handler

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitCardsHandler(postgres *gorm.DB) {
	db = postgres
}
