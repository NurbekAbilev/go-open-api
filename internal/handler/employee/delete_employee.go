package employee

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	rs := deleteEmployee(r, app.DI().EmployeeRepo)

	response.WriteJsonResponse(w, rs)
}

func deleteEmployee(r *http.Request, createEmployeeRepo repo.CreateEmployeeRepo) response.Response {
	di := app.DI()

	id := mux.Vars(r)["id"]
	if id == "" {
		return response.NewBadRequestErrorResponse("invalid id")
	}

	err := di.EmployeeRepo.DeleteEmpoyeeByID(r.Context(), id)
	if err != nil {
		return response.NewServerError(err)
	}

	return response.NewOkResponse(nil)
}
