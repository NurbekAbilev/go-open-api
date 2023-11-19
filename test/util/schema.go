package util

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nurbekabilev/go-open-api/internal/config"
	dbpkg "github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/stretchr/testify/assert"
)

func SetupSchemaForTesting(t *testing.T) (*pgxpool.Config, func()) {
	t.Helper()

	config.LoadDotEnv()
	dsn := os.Getenv("DB_URL")

	schemaName := generateSchemaNameForTest(t)

	db, err := setupConnectionForTesting(dsn)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
	assert.NoError(t, err)

	teardown := func() {
		_, err := db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schemaName))
		if err != nil {
			t.Log("teardown: could not dorp schema: %w", err)
		}
		defer db.Close()
	}

	err = dbpkg.SimpleMigrate(schemaName)
	if err != nil {
		t.Log("could not migrate: %w", err)
	}

	pgxConfig := setupPgxConfig(t, dsn, schemaName)

	return pgxConfig, teardown
}

func setupPgxConfig(t *testing.T, dsn string, schemaName string) *pgxpool.Config {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	assert.NoError(t, err)

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, fmt.Sprintf("SET search_path TO %s", schemaName))
		return err
	}

	return pgxConfig
}

func generateSchemaNameForTest(t *testing.T) string {
	schemaName := fmt.Sprintf("test_%s_%s", strings.ToLower(t.Name()), randomString(5))
	return strings.ToLower(schemaName)
}

func setupConnectionForTesting(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

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
