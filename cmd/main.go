package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/internal/handler"
)

const host = ":8080"

func main() {
	db, err := db.InitDB("postgres://postgres:example@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.InitDI(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/positions", handler.HandleAddPosition).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Listening to the host:port :", host)
	http.ListenAndServe(host, nil)
}
