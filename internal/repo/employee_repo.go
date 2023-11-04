package repo

import (
	"context"
	"database/sql"
	"time"
)

type CreateEmployeeRepo interface {
	CreateEmployee(ctx context.Context, empl Employee) (id string, err error)
}

type EmployeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) PositionRepo {
	return PositionRepo{
		db: db,
	}
}

func (r EmployeeRepo) CreateEmployee(ctx context.Context, empl Employee) (id string, err error) {
	query := `
		insert into employees(first_name, last_name, position_id, login, password, created_at) 
			values ($1, $2) 
		returning id
	`

	var lastInsertId string
	err = r.db.QueryRowContext(
		ctx, query,
		empl.FirstName, empl.LastName, empl.PositionID, empl.Login, empl.Password, time.Now(),
	).Scan(&lastInsertId)

	if err != nil {
		return "", err
	}

	return lastInsertId, nil
}
