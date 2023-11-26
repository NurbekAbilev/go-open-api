package app

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type inj struct {
	Auth         auth.AuthProvider
	PositionRepo repo.PositionRepo
	UserRepo     repo.UserRepo
	EmployeeRepo repo.EmployeeRepo
	DB           *pgxpool.Pool
}

var singleton *inj

func DI() *inj {
	if singleton == nil {
		log.Fatal("DI injector can not be nil. Initialize it with InitApp() once!")
	}
	return singleton
}

type DBInitFunc func() (*sql.DB, error)

type AppConfig struct {
	PgxCon       *pgxpool.Pool
	AuthProvider auth.AuthProvider
}

func InitApp(cfg AppConfig) error {
	if cfg.PgxCon == nil {
		log.Fatal("pgxCon cannot be nil")
	}
	if cfg.AuthProvider == nil {
		cfg.AuthProvider = auth.NewJwtAuthProvider()
	}

	initDI(cfg.PgxCon, cfg.AuthProvider)
	return nil
}

func initDI(pgxConn *pgxpool.Pool, authProvider auth.AuthProvider) {
	if pgxConn == nil {
		log.Fatal("Cannot init app with null db")
	}

	auth := auth.InitAuth()
	if authProvider != nil {
		auth = authProvider
	}

	singleton = &inj{
		EmployeeRepo: repo.NewEmployeeRepo(pgxConn),
		PositionRepo: repo.NewPositionRepo(pgxConn),
		UserRepo:     repo.NewUserRepo(pgxConn),
		Auth:         auth,
		DB:           pgxConn,
	}
}
