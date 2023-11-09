package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/auth"
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
	initDotEnv()

	dbIniter := cfg.dbIniter
	if dbIniter == nil {
		dbIniter = initDatabase
	}

	db, err := dbIniter()
	if err != nil {
		return nil, err
	}

	initDI(db)

	return func() {
		db.Close()
	}, nil
}

func initDotEnv() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase() (*sql.DB, error) {
	db, err := initDB(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
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

func initDB(connstr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
