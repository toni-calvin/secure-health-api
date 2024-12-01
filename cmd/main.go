package main

import (
	"log"
	"net/http"
	"topdoctors/db"
	"topdoctors/models"
)

func main() {
	db.Connect()

	if err := db.DB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Diagnosis{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations applied successfully.")

	db.SeedAdminUser()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
