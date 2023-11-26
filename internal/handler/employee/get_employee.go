package employee

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleGetEmployee(w http.ResponseWriter, r *http.Request) {
	rs := getEmployee(r, app.DI().EmployeeRepo)

	response.WriteJsonResponse(w, rs)
}

func getEmployee(r *http.Request, createEmployeeRepo repo.CreateEmployeeRepo) response.Response {
	di := app.DI()

	id := mux.Vars(r)["id"]
	if id == "" {
		return response.NewBadRequestErrorResponse("invalid id")
	}

	empl, err := di.EmployeeRepo.GetEmployeeById(r.Context(), id)
	if err == pgx.ErrNoRows {
		return response.NewNotFoundError("employee not found")
	}
	if err != nil {
		return response.NewServerError(err)
	}

	return response.NewOkResponse(empl)
}
