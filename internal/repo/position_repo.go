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
	CreatePosition(ctx context.Context, p Position) (id int, err error)
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

func (r PositionRepo) CreatePosition(ctx context.Context, p Position) (id int, err error) {
	query := `
		insert into positions(name, salary) 
			values ($1, $2) 
		returning id
	`

	var lastInsertId int
	err = r.db.QueryRowContext(ctx, query, p.Name, p.Salary).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("could not create position: %s", err)
	}

	return lastInsertId, nil
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
