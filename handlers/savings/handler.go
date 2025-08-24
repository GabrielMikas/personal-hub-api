package savings_handler

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitSavingsHandler(postgres *gorm.DB) {
	db = postgres
}
