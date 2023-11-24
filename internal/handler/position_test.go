package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/config"
	"github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/test/util"
)

type MockAuth struct{}
type MockUser struct{}

func (u MockUser) GetID() string {
	return "1"
}

func (u MockUser) GetRole() string {
	return "guest"
}

func (a MockAuth) GenerateToken(cred auth.Credentials) (authToken string, err error) {
	return "", nil
}
func (a MockAuth) CheckAuth(authToken string) (err error) {
	return nil
}

func (a MockAuth) GetAuthUser() auth.AbstractUser {
	return MockUser{}
}

func createMockAuth() auth.AuthProvider {
	return MockAuth{}
}

var testConn *pgxpool.Pool

func TestMain(m *testing.M) {
	config.LoadDotEnv()

	cfg, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("could not init connect")
	}

	conn, err := db.InitPgxConnect(context.TODO(), cfg)
	if err != nil {
		log.Fatalf("could not init testmain for handler: %v", err)
	}
	defer conn.Close()
	testConn = conn

	// for i := 0; i < 100; i++ {
	m.Run()
	// }

	// exitCode := m.Run()
	// os.Exit(exitCode)
}

func TestCreatePosition(t *testing.T) {
	// t.Parallel()

	tearDown := util.SetupSchemaForTesting(t, testConn)
	defer tearDown()

	err := app.InitApp(app.AppConfig{
		PgxCon:       testConn,
		AuthProvider: createMockAuth(),
	})
	assert.NoError(t, err)

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

	mux := GetRoutes()
	mux.ServeHTTP(w, r)

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

func TestGetPosition(t *testing.T) {
	// t.Parallel()

	tearDown := util.SetupSchemaForTesting(t, testConn)
	defer tearDown()

	err := app.InitApp(app.AppConfig{
		PgxCon:       testConn,
		AuthProvider: createMockAuth(),
	})
	assert.NoError(t, err)

	ctx := context.Background()

	const name = "Software Engineer"
	const salary = 1000

	var id int

	err = app.DI().DB.QueryRow(ctx, "insert into positions (name, salary) values ($1, $2) returning id", name, salary).Scan(&id)
	assert.NoError(t, err)

	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/positions/%d", id), nil)

	w := httptest.NewRecorder()

	m := GetRoutes()
	m.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusOK)

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
	assert.NoError(t, err)

	assert.Equal(t, id, rs.Data.ID, err, string(data))

}

func TestDeletePosition(t *testing.T) {
	// t.Parallel()

	tearDown := util.SetupSchemaForTesting(t, testConn)
	defer tearDown()

	err := app.InitApp(app.AppConfig{
		PgxCon:       testConn,
		AuthProvider: createMockAuth(),
	})
	assert.NoError(t, err)

	ctx := context.Background()

	var id int
	const name = "Software Engineer"
	const salary = 1000

	err = app.DI().DB.QueryRow(ctx, "insert into positions (name, salary) values ($1, $2) returning id", name, salary).Scan(&id)
	assert.NoError(t, err)

	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/positions/%d", id), nil)
	w := httptest.NewRecorder()

	m := GetRoutes()
	m.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	rs := struct {
		Code     int    `json:"code"`
		ErrorMsg string `json:"error"`
	}{}

	err = json.Unmarshal(data, &rs)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rs.Code)
	assert.Empty(t, rs.ErrorMsg)

	err = app.DI().DB.QueryRow(ctx, "select id from positions where id = $1", id).Scan(&id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("could not delete with id = %d: %v", id, err)
	}

}
