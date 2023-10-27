package app

import (
	"database/sql"

	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type inj struct {
	PositionRepo repo.PositionRepo
}

var singleton *inj

func DI() *inj {
	return singleton
}

func InitDI(db *sql.DB) {
	singleton = &inj{
		PositionRepo: repo.NewPositionRepo(db),
	}
}
