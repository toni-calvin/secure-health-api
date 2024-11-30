package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"securehealth/db"
	"securehealth/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser(t *testing.T, username, password, role string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Logf("Failed to hash password: %v", err)
	}

	err = db.TestDB.Create(&models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}).Error

	if err != nil {
		t.Logf("Failed to seed user: %v", err)
	}
	return
}

func SeedPatient(t *testing.T, patient models.Patient) {
	err := db.TestDB.Create(&patient).Error
	if err != nil {
		t.Logf("Failed to seed patient: %v", err)
	}
	return
}

func CreatePostRequest(t *testing.T, url string, body interface{}) *http.Request {
	return CreateRequest(t, "POST", url, body)
}

func CreateGetRequest(t *testing.T, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Logf("Failed to create request: %v", err)
	}
	return req
}

func CreateRequest(t *testing.T, method, url string, body interface{}) *http.Request {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Logf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		t.Logf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func ExecuteHandler(t *testing.T, handler http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	return recorder
}

func ClearTestTables() {
	tables := []string{"users", "patients", "diagnoses"}
	for _, table := range tables {
		if err := db.TestDB.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error; err != nil {
			log.Printf("Error truncating table %s: %v", table, err)
		}
	}
	fmt.Println("Tables truncated")
}

func SetupTestDatabase() {
	db.ConnectTestDatabase()
	if err := db.TestDB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Diagnosis{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	ClearTestTables()
}
