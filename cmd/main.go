package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/handler"
)

const port = "8080"

func initDb(connstr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDb("postgres://postgres:example@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	di := handler.Injector{
		DB: db,
	}

	positionHandler := handler.PositionHandler{
		DI: di,
	}

	http.Handle("/api/v1/positions/", positionHandler)

	fmt.Println("Listening port:", port)
	http.ListenAndServe(":8080", nil)
}
