package db

import (
	"log"
	"os"
	"time"
	"securehealth/constants"
	"securehealth/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SeedAdminUser() {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername == "" || adminPassword == "" {
		log.Fatalf("Admin credentials not set in environment variables. Please define [ADMIN_USERNAME, ADMIN_PASSWORD] in the .env file")
	}

	// Check if any users exist in the database
	var userCount int64
	if err := DB.Model(&models.User{}).Count(&userCount).Error; err != nil {
		log.Fatalf("Failed to check user count: %v", err)
	}

	// Seed the admin user if no users exist
	if userCount == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		adminUser := models.User{
			Username: adminUsername,
			Password: string(hashedPassword),
			Role:     constants.RoleInternal,
		}

		if err := DB.Create(&adminUser).Error; err != nil {
			log.Fatalf("Failed to seed admin user: %v", err)
		}

		log.Printf("Admin user created: username=%s, password=%s, role=%s\n", adminUsername, adminPassword, constants.RoleInternal)
	} else {
		log.Println("Admin user already exists, skipping seeding")
	}
}

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=db user=postgres password=password dbname=securehealth port=5432 sslmode=disable"
	}

	var err error
	for i := 0; i < 5; i++ { // Retry 5 times
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connected successfully.")
			return
		}
		log.Printf("Failed to connect to database. Retrying in 5 seconds... (%d/5)", i+1)
		time.Sleep(5 * time.Second)
	}

	// If all retries fail, log the fatal error
	log.Fatalf("Failed to connect to database after 5 attempts: %v", err)
}
