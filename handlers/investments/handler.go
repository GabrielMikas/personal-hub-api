package investments_handler

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitInvestmentsHandler(postgres *gorm.DB) {
	db = postgres
}
