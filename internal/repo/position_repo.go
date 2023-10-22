package repo

import (
	"database/sql"
	"errors"
	"fmt"
)

type Position struct {
	Name   *string `json:"name"`
	Salary *int    `json:"salary"`
}

type CreatePositionRepo interface {
	CreatePosition(db *sql.DB, p Position) error
}

type PositionRepo struct {
}

func (r PositionRepo) CreatePosition(db *sql.DB, p Position) error {
	_, err := db.Exec("insert into position (name, salary) values ($1, $2)", p.Name, p.Salary)
	if err != nil {
		return fmt.Errorf("could not create pos: %s", err)
	}

	return nil
}

func ValidateAddPositionStruct(p Position) error {
	if p.Name == nil || *p.Name == "" {
		return errors.New("invalid name")
	}

	if p.Salary == nil || *p.Salary < 0 {
		return errors.New("invalid salary")
	}

	return nil
}

