package repo

import (
	"context"
	"database/sql"

	"github.com/nurbekabilev/go-open-api/internal/pagination"
)

// Postition repo
type CreatePositionRepo interface {
	CreatePosition(ctx context.Context, p Position) (id int, err error)
}

type GetOnePositionRepo interface {
	GetPositionById(db *sql.DB, id int) (Position, error)
}

type GetPositionsPaginated interface {
	GetPositionsPaginated(ctx context.Context, pgReq pagination.PaginationRequest) (*pagination.PaginatedData[Position], error)
}

// Employee repo
type CreateEmployeeRepo interface {
	CreateEmployee(ctx context.Context, empl Employee) (id string, err error)
}

type GetByLoginEmployeeRepo interface {
	FindEmployeeByLogin(ctx context.Context, login string) (*Employee, error)
}
