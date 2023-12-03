package repo

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nurbekabilev/go-open-api/internal/pagination"
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

func (r EmployeeRepo) GetEmployeeById(ctx context.Context, id string) (empl Employee, err error) {
	query := `
		select id, first_name, last_name, position_id, updated_at, created_at from employees
			where id = $1
	`
	err = r.db.QueryRow(ctx, query, id).Scan(
		&empl.ID, &empl.FirstName, &empl.LastName, &empl.PositionID,
		&empl.UpdatedAt, &empl.CreatedAt,
	)
	if err != nil {
		return empl, err
	}

	return empl, nil
}

func (r EmployeeRepo) GetEmployeesPaginated(ctx context.Context, pgReq pagination.PaginationRequest) (*pagination.PaginatedData[Employee], error) {
	query := "select id, first_name, last_name, position_id, updated_at, created_at from employees"

	pageCount, rows, err := pagination.PaginateQuery(ctx, r.db, query, pgReq)
	if err != nil {
		return nil, err
	}

	empls := make([]Employee, 0)

	for rows.Next() {
		empl := Employee{}
		err := rows.Scan(
			&empl.ID, &empl.FirstName, &empl.LastName, &empl.PositionID,
			&empl.UpdatedAt, &empl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		empls = append(empls, empl)
	}

	if err != nil {
		return nil, err
	}

	pdata := pagination.PaginatedData[Employee]{
		CurentPage:    pgReq.CurrentPage,
		AmountOfPages: pageCount,
		Data:          empls,
	}

	return &pdata, nil
}

func (r EmployeeRepo) DeleteEmpoyeeByID(ctx context.Context, ID string) error {
	query := "delete from employees where id = $1"

	_, err := r.db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	return nil
}
