package users_test

import (
	"net/http"
	"testing"
	"topdoctors/api/users"
	"topdoctors/constants"
	"topdoctors/db"
	"topdoctors/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandlers(t *testing.T) {
	utils.SetupTestDatabase()

	tests := []struct {
		name           string
		setup          func(t *testing.T)
		handler        func(w http.ResponseWriter, r *http.Request)
		reqBody        users.CreateUserRequest
		expectedStatus int
	}{
		{
			name: "MissingFields",
			setup: func(t *testing.T) {
			},
			handler: users.NewUsersHandler(db.TestDB).InternalCreateUserHandler,
			reqBody: users.CreateUserRequest{
				Username: "",
				Password: "",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "ShortUsernameOrPassword",
			setup: func(t *testing.T) {
			},
			handler: users.NewUsersHandler(db.TestDB).InternalCreateUserHandler,
			reqBody: users.CreateUserRequest{
				Username: "ab",
				Password: "12345",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "DuplicateUser",
			setup: func(t *testing.T) {
				utils.SeedUser(t, "testuser", "password", constants.RoleInternal)
			},
			handler: users.NewUsersHandler(db.TestDB).InternalCreateUserHandler,
			reqBody: users.CreateUserRequest{
				Username: "testuser",
				Password: "newpassword",
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "CreateInternalUserSuccessful",
			setup: func(t *testing.T) {
			},
			handler: users.NewUsersHandler(db.TestDB).InternalCreateUserHandler,
			reqBody: users.CreateUserRequest{
				Username: "internaluser",
				Password: "securepassword",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "CreateExternalUserSuccessful",
			setup: func(t *testing.T) {
			},
			handler: users.NewUsersHandler(db.TestDB).ExternalCreateUserHandler,
			reqBody: users.CreateUserRequest{
				Username: "externaluser",
				Password: "securepassword",
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			req := utils.CreatePostRequest(t, "/users", tt.reqBody)
			recorder := utils.ExecuteHandler(t, tt.handler, req)

			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}
