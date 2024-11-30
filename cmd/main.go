package main

import (
	"fmt"
	"log"
	"net/http"
	"topdoctors/db"
	"topdoctors/models"
)

func main() {
	db.Connect()

	if err := db.DB.AutoMigrate(&models.User{}, &models.Diagnosis{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations applied successfully.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TopDoctors Challenge!")
	})
	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
