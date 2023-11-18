package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/config"
	dbpkg "github.com/nurbekabilev/go-open-api/internal/db"
)

func GenerateSchemaNameForTest(t *testing.T) string {
	schemaName := fmt.Sprintf("test_%s_%s", strings.ToLower(t.Name()), randomString(5))
	return strings.ToLower(schemaName)
}

func SetupConnectionForTesting(dsn string) (*sql.DB, error) {
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

func SetupSchemaForTesting(t *testing.T, dsn string) (string, func()) {
	t.Helper()

	schemaName := GenerateSchemaNameForTest(t)

	db, err := SetupConnectionForTesting(dsn)
	assert.NoError(t, err)

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

	return strings.ToLower(schemaName), teardown
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

func TestCreatePosition(t *testing.T) {
	t.Parallel()
	config.LoadDotEnv()

	dbUrl := os.Getenv("DB_URL")

	schemaName, tearDown := SetupSchemaForTesting(t, dbUrl)
	defer tearDown()

	pgxConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		t.Fatal(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, fmt.Sprintf("SET search_path TO %s", schemaName))
		return err
	}

	appCloser, err := app.InitApp(app.AppConfig{
		PgxConfig: pgxConfig,
	})
	assert.NoError(t, err)
	defer appCloser()

	requestBody := struct {
		Name   string `json:"name"`
		Salary int    `json:"salary"`
	}{
		Name:   "Example name",
		Salary: 1000,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("could not marshal request")
	}

	r := httptest.NewRequest(http.MethodPost, "/api/v1/positions", bytes.NewReader(requestBytes))
	w := httptest.NewRecorder()
	HandleAddPosition(w, r)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be non nil got %v", err)
	}

	rs := struct {
		Code     int    `json:"code"`
		ErrorMsg string `json:"error"`
		Data     struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Salary int    `json:"salary"`
		}
	}{}

	err = json.Unmarshal(data, &rs)
	if err != nil {
		t.Errorf("error during json unmarhsal: %v", err)
	}
	if rs.Data.Name != requestBody.Name {
		t.Fatalf("response name is not equal to request name [%s] != [%s]", rs.Data.Name, requestBody.Name)
	}

	if rs.Data.Salary != requestBody.Salary {
		t.Fatalf("response name is not equal to request name [%s] != [%s]", rs.Data.Name, requestBody.Name)
	}
}
