package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = initializePostgres()
	if err != nil {
		return fmt.Errorf("Error initializing postgres")
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
