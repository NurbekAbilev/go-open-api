package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/nurbekabilev/go-open-api/internal/pagination"
)

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

func (r PositionRepo) GetPositionsPaginated(ctx context.Context, pgReq pagination.PaginationRequest) (*pagination.PaginatedData[Position], error) {
	query := "select id, name, salary from positions"

	pageCount, rows, err := pagination.PaginateQuery(ctx, r.db, query, pgReq)
	if err != nil {
		return nil, err
	}

	positions := make([]Position, 0)

	for rows.Next() {
		pos := Position{}
		err := rows.Scan(&pos.ID, &pos.Name, &pos.Salary)
		if err != nil {
			return nil, err
		}
		positions = append(positions, pos)
	}

	if err != nil {
		log.Println("Error during find employee by login: %w", err)
		return nil, err
	}

	pdata := pagination.PaginatedData[Position]{
		CurentPage:    pgReq.CurrentPage,
		AmountOfPages: pageCount,
		Data:          positions,
	}

	return &pdata, nil
}
