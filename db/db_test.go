package db_test

import (
	"log"
	"os"
	"testing"
	"time"
	"topdoctors/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func ConnectTestDatabase() {
	dsn := "host=test_db user=postgres password=password dbname=topdoctors_test port=5432 sslmode=disable"
	var err error
	for i := 0; i < 5; i++ {
		TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Test database connected successfully.")
			return
		}
		log.Printf("Failed to connect to test database. Retrying in 5 seconds... (%d/5)", i+1)
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Failed to connect to test database after 5 attempts: %v", err)
}

func TestMain(m *testing.M) {
	ConnectTestDatabase()

	if err := TestDB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Diagnosis{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestDatabaseConnection(t *testing.T) {
	sqlDB, err := TestDB.DB()
	if err != nil {
		t.Fatalf("Error getting SQL connection: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Errorf("Error making ping to database: %v", err)
	}
}
