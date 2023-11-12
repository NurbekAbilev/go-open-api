package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nurbekabilev/go-open-api/internal/fs"
)


// Create new instance of connection(sql.DB) and migrate
func Migrate(dsn string) error {
	db, err := InitDatabase()
	if err != nil {
		return fmt.Errorf("could not init database for migration: %w", err)
	}

	mDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not init driver for migrations: %w", err)
	}

	migrationsPath := fmt.Sprintf("file://%s/migrations", fs.RootPath())
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", mDriver)
	if err != nil {
		return fmt.Errorf("could not init migration: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("could not migrate: %w", err)
	}

	return nil
}
