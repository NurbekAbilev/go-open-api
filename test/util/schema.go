package util

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	dbpkg "github.com/nurbekabilev/go-open-api/internal/db"
)

// func SetupSchemaForTesting(t *testing.T, db *sql.DB) (*pgxpool.Config, func()) {
// 	t.Helper()

// 	schemaName := generateSchemaNameForTest(t)

// 	_, err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
// 	assert.NoError(t, err)

// 	teardown := func() {
// 		_, err := db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schemaName))
// 		if err != nil {
// 			t.Log("teardown: could not dorp schema: %w", err)
// 		}
// 	}

// 	err = dbpkg.SimpleMigrate(schemaName)
// 	if err != nil {
// 		t.Log("could not migrate: %w", err)
// 	}

// 	pgxConfig := setupPgxConfig(t, os.Getenv("DB_URL"), schemaName)

// 	return pgxConfig, teardown
// }

func SetupSchemaForTesting(t *testing.T, db *pgxpool.Pool) (tearDown func()) {
	t.Helper()

	schemaName := generateSchemaNameForTest(t)
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(ctx, fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		_, err := db.Exec(ctx, fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schemaName))
		if err != nil {
			t.Log("teardown: could not drop schema: %w", err)
		}
		_ = tx.Rollback(ctx)
	}

	err = dbpkg.SimpleMigrate(ctx, db, schemaName)
	if err != nil {
		t.Logf("could not migrate schemaname = %s err: %v", schemaName, err)
	}

	_, err = db.Exec(ctx, fmt.Sprintf("SET search_path TO %s", schemaName))
	if err != nil {
		t.Fatal(err)
	}

	return teardown
}

func generateSchemaNameForTest(t *testing.T) string {
	schemaName := fmt.Sprintf("test_%s_%s", strings.ToLower(t.Name()), randomString(20))
	return strings.ToLower(schemaName)
}

// func SetupConnectionForTesting(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// Generates a random alphanumeric string of length n
func randomString(n int) string {
	const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var output []byte
	for i := 0; i < n; i++ {
		output = append(output, alphanumeric[r.Intn(len(alphanumeric))])
	}

	return string(output)
}
