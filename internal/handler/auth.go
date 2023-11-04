package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
)

func HandleAuthEmployee(w http.ResponseWriter, r *http.Request) {
	// @todo need to check login/password instead of just passing ID field
	di := app.DI()
	ctx := r.Context()

	rs := AuthorizeEmployee(ctx, w, r, di.Auth)

	response.WriteJsonResponse(w, rs)
}

func AuthorizeEmployee(ctx context.Context, w http.ResponseWriter, r *http.Request, authProvider auth.AuthProvider) response.Response {
	idStr := r.Header.Get("X-Auth-ID")

	if idStr == "" {
		return response.NewBadRequestErrorResponse("ID not provided/Invalid credentials")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequestErrorResponse("Invalid id: not integer")
	}

	cred := auth.Credentials{
		Name: "Nurbek",
		ID:   id,
	}

	token, err := authProvider.Authorize(cred)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	type TokenResponse struct {
		Token string `json:"token"`
	}

	tokenResponse := TokenResponse{
		Token: token,
	}

	return response.NewOkResponse(tokenResponse)
}
