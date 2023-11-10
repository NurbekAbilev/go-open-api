package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nurbekabilev/go-open-api/internal/fs"
)

func Migrate(db *sql.DB) error {
	migrationsPath := fmt.Sprintf("file://%s", fs.RootPath()+"/migrations")

	// migrate.New
	m, err := migrate.New(
		migrationsPath,
		os.Getenv("DB_URL"),
	)
	if err != nil {
		return fmt.Errorf("could not init migration: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("could not migrate: %w", err)
	}

	return nil
}
