package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"

	"github.com/miladibra10/vjson"
)

func TestCreatePosition(t *testing.T) {
	app.InitApp()

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

	type PositionResponse struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Salary int    `json:"salary"`
	}

	responseStruct := response.ResponseGeneric[PositionResponse]{}

	err = json.Unmarshal(data, &responseStruct)
	if err != nil {
		t.Errorf("error during json unmarhsal: %v", err)
	}

	schema := vjson.NewSchema(
		vjson.Integer("code").Required(),
		vjson.String("error").Required(),
		vjson.Object("data", vjson.NewSchema(
			vjson.Integer("id").Required(),
			vjson.String("name").Required(),
			vjson.Integer("salary1").Required(),
		)).Required(),
	)

	err = schema.ValidateBytes(data)
	if err != nil {
		t.Fatal("Error validating resposne json structure: %w", err)
	}
}
