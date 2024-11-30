package auth

import (
	"net/http"
	"testing"
	"securehealth/db"
	"securehealth/utils"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	utils.SetupTestDatabase()

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		reqBody        LoginRequest
		expectedStatus int
	}{
		{
			name: "UserNotExists",
			setup: func(t *testing.T) {
				// No user is seeded
			},
			reqBody: LoginRequest{
				Username: "nonexistent",
				Password: "password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "WrongPassword",
			setup: func(t *testing.T) {
				utils.SeedUser(t, "testuser", "correctpassword", "admin")
			},
			reqBody: LoginRequest{
				Username: "testuser",
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "LoginSuccessful",
			setup: func(t *testing.T) {
			},
			reqBody: LoginRequest{
				Username: "testuser",
				Password: "correctpassword",
			},
			expectedStatus: http.StatusOK,
		},
	}

	authHandler := NewAuthHandler(db.TestDB)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)
			req := utils.CreatePostRequest(t, "/login", tt.reqBody)
			recorder := utils.ExecuteHandler(t, authHandler.Login, req)
			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}
