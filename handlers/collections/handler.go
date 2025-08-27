package collections_handlers

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitCollectionsHandler(postgres *gorm.DB) {
	db = postgres
}
