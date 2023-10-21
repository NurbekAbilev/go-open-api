package repo

import (
	"database/sql"
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
