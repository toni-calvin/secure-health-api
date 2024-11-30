package diagnoses_test

import (
	"net/http"
	"testing"
	"securehealth/api/diagnoses"
	"securehealth/db"
	"securehealth/models"
	"securehealth/utils"

	"github.com/stretchr/testify/assert"
)

func TestDiagnosesHandler(t *testing.T) {
	utils.SetupTestDatabase()

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		method         string
		endpoint       string
		reqBody        interface{}
		expectedStatus int
	}{
		{
			name: "CreateDiagnosisPatientNotFound",
			setup: func(t *testing.T) {
			},
			method:   "POST",
			endpoint: "/diagnoses",
			reqBody: diagnoses.CreateDiagnosisRequest{
				PatientID: "123e4567-e89b-12d3-a456-426614174000",
				Diagnosis: "Hypertension",
				StartDate: "2023-01-01",
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "CreateDiagnosisSuccessful",
			setup: func(t *testing.T) {
				utils.SeedPatient(t, models.Patient{
					ID:   "123e4567-e89b-12d3-a456-426614174000",
					Name: "testpatient",
				})
			},
			method:   "POST",
			endpoint: "/diagnoses",
			reqBody: diagnoses.CreateDiagnosisRequest{
				PatientID:    "123e4567-e89b-12d3-a456-426614174000",
				Diagnosis:    "Hypertension",
				Prescription: "Medication A",
				StartDate:    "2023-01-01",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "GetDiagnosesSuccessful",
			setup: func(t *testing.T) {
			},
			method:         "GET",
			endpoint:       "/diagnoses?name=testpatient",
			reqBody:        nil,
			expectedStatus: http.StatusOK,
		},
	}

	diagnosesHandler := diagnoses.NewDiagnosesHandler(db.TestDB)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			req := utils.CreateRequest(t, tt.method, tt.endpoint, tt.reqBody)
			recorder := utils.ExecuteHandler(t, func(w http.ResponseWriter, r *http.Request) {
				if tt.method == "POST" {
					diagnosesHandler.CreateDiagnosis(w, r)
				} else {
					diagnosesHandler.GetDiagnoses(w, r)
				}
			}, req)

			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}
