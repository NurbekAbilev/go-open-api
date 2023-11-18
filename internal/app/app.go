package app

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	dbpkg "github.com/nurbekabilev/go-open-api/internal/db"
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
	PgxConfig  *pgxpool.Config
	SchemaName string
}

func InitApp(cfg AppConfig) (closer func(), err error) {
	ctx := context.Background()
	dbUrl := os.Getenv("DB_URL")

	if cfg.PgxConfig == nil {
		conf, err := pgxpool.ParseConfig(dbUrl)
		if err != nil {
			return nil, err
		}
		cfg.PgxConfig = conf
	}

	pgxConn, err := dbpkg.InitPgxConnect(ctx, cfg.PgxConfig)
	if err != nil {
		return nil, err
	}

	initDI(pgxConn)

	return func() {
		pgxConn.Close()
	}, nil
}

func initDI(pgxConn *pgxpool.Pool) {
	if pgxConn == nil {
		log.Fatal("Cannot init app with null db")
	}

	singleton = &inj{
		// EmployeeRepo: repo.NewEmployeeRepo(db),
		PositionRepo: repo.NewPositionRepo(pgxConn),
		UserRepo:     repo.NewUserRepo(pgxConn),
		Auth:         auth.InitAuth(),
		DB:           pgxConn,
	}
}
