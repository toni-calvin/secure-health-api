package api

import (
	"topdoctors/api/auth"
	"topdoctors/api/users"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Login Endpoint
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	// Internal API
	internal := r.PathPrefix("/internal").Subrouter()
	internal.Use(auth.AuthMiddleware("internal"))
	internal.HandleFunc("/users", users.InternalCreateUserHandler).Methods("POST")

	// internal.HandleFunc("/diagnoses", diagnoses.GetDiagnoses).Methods("GET")
	// internal.HandleFunc("/diagnoses", diagnoses.CreateDiagnosis).Methods("POST")

	// External API
	external := r.PathPrefix("/external").Subrouter()
	external.Use(auth.AuthMiddleware("external"))
	external.HandleFunc("/users", users.InternalCreateUserHandler).Methods("POST")

	// external.HandleFunc("/diagnoses", diagnoses.GetDiagnoses).Methods("GET")
	// external.HandleFunc("/diagnoses", diagnoses.CreateDiagnosis).Methods("POST")

	return r
}
