package patients

import (
	"encoding/json"
	"net/http"
	"strings"
	"topdoctors/db"
	"topdoctors/models"
)

type CreatePatientRequest struct {
	Name    string `json:"name"`
	NIF     string `json:"nif"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePatientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Name == "" || req.NIF == "" || req.Email == "" {
		http.Error(w, "Name, NIF, and Email are required fields", http.StatusBadRequest)
		return
	}

	patient := models.Patient{
		Name:    req.Name,
		NIF:     strings.TrimSpace(req.NIF),
		Email:   strings.TrimSpace(req.Email),
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := db.DB.Create(&patient).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			http.Error(w, "Patient with the same NIF or Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

func ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")

	var patients []models.Patient
	query := db.DB

	if nameFilter != "" {
		query = query.Where("LOWER(name) = ?", strings.ToLower(nameFilter))
	}

	if err := query.Find(&patients).Error; err != nil {
		http.Error(w, "Failed to fetch patients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}
