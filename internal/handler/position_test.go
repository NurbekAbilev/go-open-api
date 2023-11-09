package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"

	"github.com/nurbekabilev/go-open-api/internal/app"

	"github.com/miladibra10/vjson"
)

func createSchema(db *sql.DB, schemaName string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE SCHEMA %s", schemaName))
	return err
}

func dropSchema(db *sql.DB, schemaName string) error {
	_, err := db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schemaName))
	return err
}

func setupSchema(t *testing.T, db *sql.DB) (string, func()) {
	t.Helper()

	schemaName := fmt.Sprintf("test_schema_%s", t.Name())

	err := createSchema(db, schemaName)
	if err != nil {
		t.Fatalf("could not create schema: %v", err)
	}

	teardown := func() {
		if err := dropSchema(db, schemaName); err != nil {
			t.Fatalf("could not drop schema: %v", err)
		}
	}

	return schemaName, teardown
}

func TestCreatePosition(t *testing.T) {
	t.Parallel()

	close, err := app.InitApp(app.AppConfig{})
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	

	// app.DI().

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

	type posResponse struct {
		ID   int
		Name int
	}

	type responseStruct struct {
		Code     int    `json:"code"`
		ErrorMsg string `json:"error"`
		Data     struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Salary int    `json:"salary"`
		}
	}

	rs := responseStruct{}

	err = json.Unmarshal(data, &rs)
	if err != nil {
		t.Errorf("error during json unmarhsal: %v", err)
	}

	schema := vjson.NewSchema(
		vjson.Integer("code").Required(),
		vjson.String("error").Required(),
		vjson.Object("data", vjson.NewSchema(
			vjson.Integer("id").Required(),
			vjson.String("name").Required(),
			vjson.Integer("salary").Required(),
		)).Required(),
	)

	if rs.Data.Name != requestBody.Name {
		t.Fatalf("response name is not equal to request name [%s] != [%s]", rs.Data.Name, requestBody.Name)
	}

	if rs.Data.Salary != requestBody.Salary {
		t.Fatalf("response name is not equal to request name [%s] != [%s]", rs.Data.Name, requestBody.Name)
	}

	err = schema.ValidateBytes(data)
	if err != nil {
		t.Fatal("Error validating resposne json structure: %w", err)
	}
}
