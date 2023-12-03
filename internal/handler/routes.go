package handler

import (
	"github.com/gorilla/mux"
	"github.com/nurbekabilev/go-open-api/internal/handler/employee"
	"github.com/nurbekabilev/go-open-api/internal/middleware"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Auth routes
	r.HandleFunc("/api/v1/signup", HandleSignUp).Methods("POST")
	r.HandleFunc("/api/v1/auth", HandleAuth).Methods("POST")
	
	// Positions
	protectedRoutes.HandleFunc("/api/v1/positions", HandleGetPositions).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/positions/{id:[0-9]+}", HandleGetOnePosition).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/positions", HandleAddPosition).Methods("POST")
	protectedRoutes.HandleFunc("/api/v1/positions/{id:[0-9]+}", HandleDeletePosition).Methods("DELETE")

	// Employees 
	// @todo finish swagger
	protectedRoutes.HandleFunc("/api/v1/employees", employee.HandleGetEmployeeList).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/employees/{id:[0-9]+}", employee.HandleGetEmployee).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/employees", employee.HandleAddEmployee).Methods("POST")
	protectedRoutes.HandleFunc("/api/v1/employees/{id:[0-9]+}", employee.HandleDeleteEmployee).Methods("DELETE")

	return r
}
