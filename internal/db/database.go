package db

import (
	"database/sql"
	"errors"
	"os"
)

func InitDatabase() (*sql.DB, error) {
	connstr := os.Getenv("DB_URL")
	if connstr == "" {
		return nil, errors.New("empty DB_URL variable, set inside .env")
	}

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
