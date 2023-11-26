package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/pagination"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleAddPosition(w http.ResponseWriter, r *http.Request) {
	di := app.DI()
	rs := addPosition(r, di.PositionRepo)

	response.WriteJsonResponse(w, rs)
}

func HandleGetPositions(w http.ResponseWriter, r *http.Request) {
	// di := app.DI()
	ctx := r.Context()

	rs := getPositions(ctx, w, r)

	response.WriteJsonResponse(w, rs)
}

func HandleGetOnePosition(w http.ResponseWriter, r *http.Request) {
	rs := getOnePositions(r.Context(), r)
	response.WriteJsonResponse(w, rs)
}

func HandleDeletePosition(w http.ResponseWriter, r *http.Request) {
	rs := deleteOnePosition(r.Context(), r)
	response.WriteJsonResponse(w, rs)
}

func getOnePositions(ctx context.Context, r *http.Request) response.Response {
	di := app.DI()

	id := mux.Vars(r)["id"]
	if id == "" {
		return response.NewBadRequestErrorResponse("invalid id")
	}

	pos, err := di.PositionRepo.GetOnePositionByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return response.NewNotFoundError("position not found")
	}
	if err != nil {
		return response.NewServerError(err)
	}

	return response.NewOkResponse(pos)
}

func getPositions(ctx context.Context, w http.ResponseWriter, r *http.Request) response.Response {
	di := app.DI()

	rg := pagination.PaginationRequest{
		PerPageAmount: pagination.LimitPerPage,
		CurrentPage:   2,
	}

	pd, err := di.PositionRepo.GetPositionsPaginated(ctx, rg)
	if err != nil {
		return response.NewServerError(err)
	}

	return response.NewOkResponse(pd)
}

func addPosition(r *http.Request, createPosRepo repo.CreatePositionRepo) response.Response {
	p := repo.Position{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("could not decode positoin: %v", err)
		return response.NewServerError(err)
	}

	err = repo.ValidateAddPositionStruct(p)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	id, err := createPosRepo.CreatePosition(r.Context(), p)
	if err != nil {
		log.Printf("could not create position: %v", err)
		return response.NewServerError(err)
	}

	p.ID = &id

	return response.NewOkResponse(p)
}

func deleteOnePosition(ctx context.Context, r *http.Request) response.Response {
	di := app.DI()

	id := mux.Vars(r)["id"]
	if id == "" {
		return response.NewBadRequestErrorResponse("invalid id")
	}

	err := di.PositionRepo.DeleteOnePositionByID(ctx, id)
	if err != nil {
		response.NewServerError(err)
	}

	return response.NewOkResponse(nil)
}
