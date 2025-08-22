package config

import (
	"fmt"
	"os"

	"github.com/GabrielMikas/personal-hub-api/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		fmt.Println("Error: DB_DSN not set in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening database:", err)
	}

	err = db.AutoMigrate(&schemas.Investment{})
	if err != nil {
		fmt.Println("Error running migration:", err)
	}

	return db, nil
}
