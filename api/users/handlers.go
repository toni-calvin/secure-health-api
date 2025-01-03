package users

import (
	"encoding/json"
	"log"
	"net/http"
	"securehealth/constants"
	"securehealth/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersHandler struct {
	DB *gorm.DB
}

func NewUsersHandler(db *gorm.DB) *UsersHandler {
	return &UsersHandler{DB: db}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

func jsonError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *UsersHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request, role string) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(req.Username) < 3 || len(req.Password) < 6 {
		jsonError(w, "Username must be at least 3 characters and password at least 6 characters long", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	if err := h.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		jsonError(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		jsonError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     role,
	}
	if err := h.DB.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
		jsonError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{Message: "User created successfully"})
}

func (h *UsersHandler) InternalCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	h.CreateUserHandler(w, r, constants.RoleInternal)
}

func (h *UsersHandler) ExternalCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	h.CreateUserHandler(w, r, constants.RoleExternal)
}
