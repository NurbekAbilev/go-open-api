package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/internal/handler"
)

const host = ":8080"

func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(path.Dir(b)))
}

func main() {
	rootPath := getProjectRoot()
	
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.InitDB(os.Getenv("DB_URL"))
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
