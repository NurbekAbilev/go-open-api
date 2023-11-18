package main

import (
	"fmt"
	"log"
	"net/http"

	sw "github.com/flowchartsman/swaggerui"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/config"
	"github.com/nurbekabilev/go-open-api/internal/handler"
	"github.com/nurbekabilev/go-open-api/internal/middleware"

	"github.com/nurbekabilev/go-open-api/internal/handler/swaggerui"
)

const host = ":8080"

func main() {
	config.LoadDotEnv()

	// todo migrate before init app

	closer, err := app.InitApp(app.AppConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	r := mux.NewRouter()
	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Positions
	protectedRoutes.HandleFunc("/api/v1/positions", handler.HandleAddPosition).Methods("POST")
	protectedRoutes.HandleFunc("/api/v1/positions", handler.HandleGetPositions).Methods("GET")
	protectedRoutes.HandleFunc("/api/v1/positions/{id:[0-9]+}", handler.HandleGetOnePosition).Methods("GET")

	// Employees @todo

	// Auth routes
	r.HandleFunc("/api/v1/signup", handler.HandleSignUp).Methods("POST")
	r.HandleFunc("/api/v1/auth", handler.HandleAuth).Methods("POST")
	// r.HandleFunc("/api/v1/auth/validate", handler.ValidateAuthEmployee).Methods("POST")

	// Host swagger-ui
	http.Handle("/api/swagger/", http.StripPrefix("/api/swagger", sw.Handler(swaggerui.GetSwaggerYml())))

	// Setup handler
	http.Handle("/", r)

	fmt.Println("Listening to the:", host)
	err = http.ListenAndServe(host, nil)
	if err != nil {
		panic(err)
	}
}
