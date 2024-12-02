package api

import (
	"topdoctors/api/auth"
	"topdoctors/api/diagnoses"
	"topdoctors/api/patients"
	"topdoctors/api/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	authHandler := auth.NewAuthHandler(db)
	usersHandler := users.NewUsersHandler(db)
	diagnosesHandler := diagnoses.NewDiagnosesHandler(db)
	patientsHandler := patients.NewPatientsHandler(db)

	// Login Endpoint
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Internal API
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(auth.AuthMiddleware("internal"))
	internal.HandleFunc("/users", usersHandler.InternalCreateUserHandler).Methods("POST")
	internal.HandleFunc("/patients", patientsHandler.CreatePatientHandler).Methods("POST")
	internal.HandleFunc("/patients", patientsHandler.ListPatientsHandler).Methods("GET")
	internal.HandleFunc("/diagnoses", diagnosesHandler.GetDiagnoses).Methods("GET")
	internal.HandleFunc("/diagnoses", diagnosesHandler.CreateDiagnosis).Methods("POST")

	// External API
	external := r.PathPrefix("/external").Subrouter()
	external.Use(auth.AuthMiddleware("external"))
	external.HandleFunc("/users", usersHandler.ExternalCreateUserHandler).Methods("POST")
	external.HandleFunc("/patients", patientsHandler.CreatePatientHandler).Methods("POST")
	external.HandleFunc("/patients", patientsHandler.ListPatientsHandler).Methods("GET")
	external.HandleFunc("/diagnoses", diagnosesHandler.GetDiagnoses).Methods("GET")
	external.HandleFunc("/diagnoses", diagnosesHandler.CreateDiagnosis).Methods("POST")

	return r
}
