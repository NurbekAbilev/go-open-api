package handler

import (
	"github.com/gorilla/mux"
	"github.com/nurbekabilev/go-open-api/internal/middleware"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Positions
	protectedRoutes.HandleFunc("/api/v1/positions", HandleAddPosition).Methods("POST")
	protectedRoutes.HandleFunc("/api/v1/positions", HandleGetPositions).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/positions/{id:[0-9]+}", HandleGetOnePosition).Methods("GET")

	// Employees @todo

	// Auth routes
	r.HandleFunc("/api/v1/signup", HandleSignUp).Methods("POST")
	r.HandleFunc("/api/v1/auth", HandleAuth).Methods("POST")

	return r
}
