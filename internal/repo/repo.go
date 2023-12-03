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

type GetOnePositionByID interface {
	GetOnePositionByID(ctx context.Context, ID string) (Position, error)
}

type DeletePositionByIdRepo interface {
	DeleteOnePositionByID(ctx context.Context, ID string) error
}

// Employee repo
type CreateEmployeeRepo interface {
	CreateEmployee(ctx context.Context, empl Employee) (id string, err error)
}

type GetOneEmployeeByIDRepo interface {
	GetEmployeeById(ctx context.Context, id string) (empl Employee, err error)
}

type GetEmployeesPaginated interface {
	GetEmployeesPaginated(ctx context.Context, pgReq pagination.PaginationRequest) (*pagination.PaginatedData[Employee], error)
}

type DeleteEmployeeByIDRepo interface {
	DeleteEmpoyeeByID(ctx context.Context, ID string) error
}

// User repo
type CreateUserRepo interface {
	CreateUser(ctx context.Context, user User) (string, error)
}

type GetUserByLoginRepo interface {
	GetUserByLogin(ctx context.Context, login string) (*User, error)
}
