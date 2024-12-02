package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var TestDB *gorm.DB

func ConnectTestDatabase() {
	dsn := "host=test_db user=postgres password=password dbname=topdoctors_test port=5432 sslmode=disable"
	var err error
	for i := 0; i < 5; i++ {
		TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // Suppress logs
		})
		if err == nil {
			return
		}
		log.Printf("Failed to connect to test database. Retrying in 5 seconds... (%d/5)", i+1)
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Failed to connect to test database after 5 attempts: %v", err)
}
