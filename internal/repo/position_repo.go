package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Position struct {
	ID     *int    `json:"id"`
	Name   *string `json:"name"`
	Salary *int    `json:"salary"`
}

type CreatePositionRepo interface {
	CreatePosition(ctx context.Context, p Position) error
}

type GetOnePositionRepo interface {
	GetPositionById(db *sql.DB, id int) (Position, error)
}

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) PositionRepo {
	return PositionRepo{
		db: db,
	}
}

func (r PositionRepo) CreatePosition(ctx context.Context, p Position) error {
	_, err := r.db.ExecContext(ctx, "insert into position (name, salary) values ($1, $2)", p.Name, p.Salary)
	if err != nil {
		return fmt.Errorf("could not create pos: %s", err)
	}

	return nil
}

// func (repo PositionRepo) GetPositionById(db *sql.DB, id int) (*Position, error) {
// 	var p Position

// 	row := db.QueryRow("select id, name, salary from position where id = $1", id)
// 	if err := row.Scan(&p.ID, &p.Name, &p.Salary); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &p, nil
// }

func ValidateAddPositionStruct(p Position) error {
	if p.Name == nil || *p.Name == "" {
		return errors.New("invalid name")
	}

	if p.Salary == nil || *p.Salary < 0 {
		return errors.New("invalid salary")
	}

	return nil
}
