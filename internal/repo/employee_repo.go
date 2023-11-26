package repo

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) EmployeeRepo {
	return EmployeeRepo{
		db: db,
	}
}

func (r EmployeeRepo) CreateEmployee(ctx context.Context, empl Employee) (id string, err error) {
	query := `
		insert into employees(first_name, last_name, position_id, created_at) 
			values ($1, $2, $3, $4) 
		returning id
	`

	var lastInsertId string
	err = r.db.QueryRow(
		ctx, query,
		empl.FirstName, empl.LastName, empl.PositionID, time.Now(),
	).Scan(&lastInsertId)

	if err != nil {
		log.Println("CreateEmployee: %w", err)
		return "", err
	}

	return lastInsertId, nil
}
