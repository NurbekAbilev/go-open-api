//go
package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/config"
	"github.com/nurbekabilev/go-open-api/test/util"
)

func TestCreatePosition(t *testing.T) {
	t.Parallel()
	config.LoadDotEnv()

	dbUrl := os.Getenv("DB_URL")

	pgxConfig, tearDown := util.SetupSchemaForTesting(t, dbUrl)
	defer tearDown()

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
