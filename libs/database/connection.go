package database

import (
	"example/libs/database/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	driver := os.Getenv("DB_DRIVER")

	if driver == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Falha ao conectar-se à base de dados!")
	}

	DB.AutoMigrate(&models.User{}, &models.Product{})

	return nil
}
