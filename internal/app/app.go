package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/fs"
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
	db := InitDatabase()
	InitDI(db)

	return func() {
		log.Println("Clsoing ....")
		db.Close()
	}
}

func InitConfig() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func InitDatabase() (db *sql.DB) {
	db, err := initDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
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
