package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default fallback for Docker Compose environment
		dsn = "host=db user=postgres password=password dbname=topdoctors port=5432 sslmode=disable"
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
