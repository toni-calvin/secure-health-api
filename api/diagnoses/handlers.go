package diagnoses

import (
	"encoding/json"
	"net/http"
	"time"
	"securehealth/models"

	"gorm.io/gorm"
)

type DiagnosesHandler struct {
	DB *gorm.DB
}

func NewDiagnosesHandler(db *gorm.DB) *DiagnosesHandler {
	return &DiagnosesHandler{DB: db}
}

type CreateDiagnosisRequest struct {
	PatientID    string `json:"patient_id"`
	Diagnosis    string `json:"diagnosis"`
	Prescription string `json:"prescription"`
	StartDate    string `json:"start_date"`
}

func (h *DiagnosesHandler) GetDiagnoses(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	startDate := r.URL.Query().Get("start_date")

	var diagnoses []models.Diagnosis
	query := h.DB.Table("diagnoses").Joins("INNER JOIN patients ON patients.id = diagnoses.patient_id")

	if name != "" {
		query = query.Where("patients.name ILIKE ?", "%"+name+"%")
	}
	if startDate != "" {
		query = query.Where("DATE(diagnoses.start_date) = ?", startDate)
	}

	if err := query.Find(&diagnoses).Error; err != nil {
		http.Error(w, "Failed to fetch diagnoses", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(diagnoses)
}

func (h *DiagnosesHandler) CreateDiagnosis(w http.ResponseWriter, r *http.Request) {
	var req CreateDiagnosisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.PatientID == "" || req.Diagnosis == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var patient models.Patient
	if err := h.DB.Where("id = ?", req.PatientID).First(&patient).Error; err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	var existingDiagnosis models.Diagnosis
	if err := h.DB.Where("patient_id = ? AND diagnosis = ?", req.PatientID, req.Diagnosis).First(&existingDiagnosis).Error; err == nil {
		http.Error(w, "Diagnosis already exists for this patient", http.StatusConflict)
		return
	}

	layout := "2006-01-02" // Example format: YYYY-MM-DD
	parsedDate, err := time.Parse(layout, req.StartDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	diagnose := models.Diagnosis{
		PatientID:    req.PatientID,
		Prescription: req.Prescription,
		Diagnosis:    req.Diagnosis,
		StartDate:    parsedDate,
	}

	if err := h.DB.Create(&diagnose).Error; err != nil {
		http.Error(w, "Failed to create diagnosis", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
