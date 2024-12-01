package api

import (
	"topdoctors/api/auth"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Login Endpoint
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	return r
}
