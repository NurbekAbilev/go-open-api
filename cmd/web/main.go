package main

import (
	"fmt"
	"net/http"

	"github.com/flowchartsman/swaggerui"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler"
	"github.com/nurbekabilev/go-open-api/internal/middleware"

	sw "github.com/nurbekabilev/go-open-api/internal/handler/swaggerui"
)

const host = ":8080"

func main() {
	closer := app.InitApp()
	defer closer()

	r := mux.NewRouter()
	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Positions
	protectedRoutes.HandleFunc("/api/v1/positions", handler.HandleAddPosition).Methods("POST")

	// Employees @todo

	// Auth routes
	r.HandleFunc("/api/v1/signup", handler.HandleAuthEmployee).Methods("POST")
	r.HandleFunc("/api/v1/auth", handler.HandleAuthEmployee).Methods("POST")
	// r.HandleFunc("/api/v1/auth/validate", handler.ValidateAuthEmployee).Methods("POST")

	// Host swagger-ui
	http.Handle("/api/swagger/", http.StripPrefix("/api/swagger", swaggerui.Handler(sw.GetSwaggerYml())))

	// Setup handler
	http.Handle("/", r)

	fmt.Println("Listening to the:", host)
	http.ListenAndServe(host, nil)
}
