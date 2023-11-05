package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/pagination"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleAddPosition(w http.ResponseWriter, r *http.Request) {
	di := app.DI()
	ctx := r.Context()

	rs := addPosition(ctx, di.PositionRepo, r)

	response.WriteJsonResponse(w, rs)
}

func HandlerGetPositions(w http.ResponseWriter, r *http.Request) {
	// di := app.DI()
	ctx := r.Context()

	rs := getPositions(ctx, w, r)

	response.WriteJsonResponse(w, rs)
}

func getPositions(ctx context.Context, w http.ResponseWriter, r *http.Request) response.Response {
	di := app.DI()

	rg := pagination.PaginationRequest{
		PerPageAmount: pagination.LimitPerPage,
		CurrentPage:   2,
	}

	pd, err := di.PositionRepo.GetPositionsPaginated(ctx, rg)
	if err != nil {
		return response.NewServerError("server error")
	}

	return response.NewOkResponse(pd)
}

func addPosition(ctx context.Context, createPosRepo repo.CreatePositionRepo, r *http.Request) response.Response {
	p := repo.Position{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("could not decode positoin: %v", err)
		return response.NewServerError(err.Error())
	}

	err = repo.ValidateAddPositionStruct(p)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	id, err := createPosRepo.CreatePosition(ctx, p)
	if err != nil {
		log.Printf("could not create position: %v", err)
		return response.NewServerError(err.Error())
	}

	p.ID = &id

	return response.NewOkResponse(p)
}
