package app

import (
	"database/sql"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type inj struct {
	PositionRepo repo.PositionRepo
}

var singleton *inj

func DI() *inj {
	return singleton
}

func InitApp() (closer func()) {
	InitConfig()

	db, dbCloser := InitDatabase()

	InitDI(db)

	return func() {
		dbCloser()
	}
}

func InitConfig() {
	rootPath := rootPath() + "/../.."
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func InitDatabase() (db *sql.DB, closer func()) {
	db, err := initDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db, func() { db.Close() }
}

func InitDI(db *sql.DB) {
	singleton = &inj{
		PositionRepo: repo.NewPositionRepo(db),
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

func rootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(b))
}
