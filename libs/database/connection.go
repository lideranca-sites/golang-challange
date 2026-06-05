package database

import (
	"example/libs/database/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	log.Println("Attempting to connect to database...")
	log.Printf("Host: %s, DB: %s, User: %s\n", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully!")

	// Try to migrate, if it fails due to type conversion issues, drop and recreate
	err = DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		// If migration fails, likely due to old string timestamp columns
		// Drop tables and recreate with correct schema
		log.Println("Migration failed, dropping and recreating tables...")
		DB.Migrator().DropTable(&models.User{}, &models.Product{})

		// Retry migration after dropping tables
		err = DB.AutoMigrate(&models.User{}, &models.Product{})
		if err != nil {
			return fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	log.Println("Migrations completed successfully!")
	return nil
}
