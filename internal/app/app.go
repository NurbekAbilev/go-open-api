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
}

var singleton *inj

func DI() *inj {
	if singleton == nil {
		log.Fatal("DI injector can not be nil. Initialize it with InitApp() once!")
	}
	return singleton
}

func InitApp() (closer func()) {
	initConfig()
	db := initDatabase()
	initDI(db)

	return func() {
		db.Close()
	}
}

func initConfig() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase() (db *sql.DB) {
	db, err := initDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initDI(db *sql.DB) {
	singleton = &inj{
		PositionRepo: repo.NewPositionRepo(db),
		Auth:         auth.InitAuth(),
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
