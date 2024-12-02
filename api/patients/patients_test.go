package patients_test

import (
	"net/http"
	"testing"
	"topdoctors/api/patients"
	"topdoctors/db"
	"topdoctors/models"
	"topdoctors/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreatePatientHandler(t *testing.T) {
	utils.SetupTestDatabase()

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		reqBody        patients.CreatePatientRequest
		expectedStatus int
	}{
		{
			name: "MissingRequiredFields",
			setup: func(t *testing.T) {
			},
			reqBody: patients.CreatePatientRequest{
				Name:  "",
				NIF:   "123456789",
				Email: "",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "DuplicatePatient",
			setup: func(t *testing.T) {
				utils.SeedPatient(t, models.Patient{
					Name:  "testpatient",
					NIF:   "123456789",
					Email: "testpatient@example.com",
				})
			},
			reqBody: patients.CreatePatientRequest{
				Name:  "testpatient",
				NIF:   "123456789",
				Email: "testpatient@example.com",
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "SuccessfulPatientCreation",
			setup: func(t *testing.T) {
			},
			reqBody: patients.CreatePatientRequest{
				Name:    "testpatient2",
				NIF:     "987654321",
				Email:   "testpatient2@example.com",
				Phone:   "1234567890",
				Address: "address",
			},
			expectedStatus: http.StatusCreated,
		},
	}

	handler := patients.NewPatientsHandler(db.TestDB)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			req := utils.CreatePostRequest(t, "/patients", tt.reqBody)
			recorder := utils.ExecuteHandler(t, handler.CreatePatientHandler, req)

			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}
