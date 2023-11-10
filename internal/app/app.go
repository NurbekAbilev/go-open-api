package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/db"
	dbpkg "github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/internal/fs"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type inj struct {
	Auth         auth.AuthProvider
	PositionRepo repo.PositionRepo
	EmployeeRepo repo.EmployeeRepo
	DB           *sql.DB
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
	dbIniter DBInitFunc
}

func InitApp(cfg AppConfig) (closer func(), err error) {
	loadDotEnv()

	dbIniter := cfg.dbIniter
	if dbIniter == nil {
		dbIniter = db.InitDatabase
	}

	db, err := dbIniter()
	if err != nil {
		return nil, err
	}

	initDI(db)
	err = dbpkg.Migrate(db)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("could not migrate: %w", err)
	}

	return func() {
		db.Close()
	}, nil
}

func loadDotEnv() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func initDI(db *sql.DB) {
	if db == nil {
		log.Fatal("Cannot init app with null db")
	}

	singleton = &inj{
		EmployeeRepo: repo.NewEmployeeRepo(db),
		PositionRepo: repo.NewPositionRepo(db),
		Auth:         auth.InitAuth(),
		DB:           db,
	}
}
