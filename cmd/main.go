package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/internal/handler"
)

const host = ":8080"

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.Getenv("JWT_KEY"))

	db, err := db.InitDB("postgres://postgres:example@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.InitDI(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/positions", handler.HandleAddPosition).Methods("POST")
	r.HandleFunc("/api/v1/auth", handler.HandleAuthEmployee).Methods("POST")
	r.HandleFunc("/api/v1/auth/validate", handler.ValidateAuthEmployee).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Listening to the:", host)
	http.ListenAndServe(host, nil)
}
