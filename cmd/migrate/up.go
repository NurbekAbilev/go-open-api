package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	m, err := migrate.New(
		"file:///migrations",
		os.Getenv("DB_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
