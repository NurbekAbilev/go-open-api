package employee

import (
	"encoding/json"
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleAddEmployee(w http.ResponseWriter, r *http.Request) {
	rs := addEmployee(r, app.DI().EmployeeRepo)

	response.WriteJsonResponse(w, rs)
}

func addEmployee(r *http.Request, createEmployeeRepo repo.CreateEmployeeRepo) response.Response {
	empl := repo.Employee{}
	err := json.NewDecoder(r.Body).Decode(&empl)
	if err != nil {
		return response.NewServerError(err)
	}

	id, err := createEmployeeRepo.CreateEmployee(r.Context(), empl)
	if err != nil {
		return response.NewServerError(err)
	}
	// @todo remove hardcode
	// empl.ID = id
	_ = id
	empl.ID = 1

	return response.NewOkResponse(empl)
}
