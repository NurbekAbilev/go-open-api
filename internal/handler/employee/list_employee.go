package employee

import (
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/pagination"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleGetEmployeeList(w http.ResponseWriter, r *http.Request) {
	rs := getEmployeeList(r, app.DI().EmployeeRepo)

	response.WriteJsonResponse(w, rs)
}

func getEmployeeList(r *http.Request, createEmployeeRepo repo.CreateEmployeeRepo) response.Response {
	di := app.DI()

	rg := pagination.PaginationRequest{
		PerPageAmount: pagination.LimitPerPage,
		CurrentPage:   1,
	}

	pd, err := di.EmployeeRepo.GetEmployeesPaginated(r.Context(), rg)
	if err != nil {
		return response.NewServerError(err)
	}

	return response.NewOkResponse(pd)
}
