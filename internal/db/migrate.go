package db

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	rtHelper "github.com/nurbekabilev/go-open-api/internal/fs"
)

// Create new instance of connection(sql.DB) and migrate
func Migrate(dsn string) error {
	if dsn == "" {
		return errors.New("empty dsn: set DB_URL variable")
	}

	// db, err := InitDatabase()
	// if err != nil {
	// 	return fmt.Errorf("could not init database for migration: %w", err)
	// }
	// _, err = db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName))
	// if err != nil {
	// 	return fmt.Errorf("could not init database for migration: %w", err)
	// }
	migrationsPath := fmt.Sprintf("file://%s/migrations", rtHelper.RootPath())

	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		return fmt.Errorf("could not init migration: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("could not migrate: %w", err)
	}

	return nil
}

func getUpMigrationsFiles(migrationsPath string) []string {
	var files []string

	filepath.WalkDir(migrationsPath, func(s string, d fs.DirEntry, e error) error {
		if strings.HasSuffix(s, ".up.sql") {
			files = append(files, s)
		}
		return nil
	})

	return files
}

func SimpleMigrate(schemaName string) error {

	migrationsPath := rtHelper.RootPath() + "/migrations"

	files := getUpMigrationsFiles(migrationsPath)

	db, err := InitDatabase()
	if err != nil {
		return fmt.Errorf("could not init db: %w", err)
	}

	schemaName = strings.ToLower(schemaName)

	qr := fmt.Sprintf(`SET search_path TO %s`, schemaName)
	_, err = db.Exec(qr)
	if err != nil {
		return err
	}

	for _, file := range files {
		contents, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		migrationsQuery := string(contents)

		_, err = db.Exec(migrationsQuery)
		if err != nil {
			return err
		}

		// log.Printf("migrated %s", file)
	}

	return nil
}
