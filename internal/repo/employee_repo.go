package repo

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type EmployeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) EmployeeRepo {
	return EmployeeRepo{
		db: db,
	}
}

func (r EmployeeRepo) CreateEmployee(ctx context.Context, empl Employee) (id string, err error) {
	query := `
		insert into employees(first_name, last_name, position_id, login, password, created_at) 
			values ($1, $2, $3, $4, $5, $6) 
		returning id
	`

	var lastInsertId string
	err = r.db.QueryRowContext(
		ctx, query,
		empl.FirstName, empl.LastName, empl.PositionID, empl.Login, empl.Password, time.Now(),
	).Scan(&lastInsertId)

	if err != nil {
		log.Println("CreateEmployee: %w", err)
		return "", err
	}

	return lastInsertId, nil
}