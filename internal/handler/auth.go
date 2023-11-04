package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleAuthEmployee(w http.ResponseWriter, r *http.Request) {
	// @todo need to check login/password instead of just passing ID field
	di := app.DI()
	ctx := r.Context()

	rs := AuthorizeEmployee(ctx, w, r, di.Auth)

	response.WriteJsonResponse(w, rs)
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rs := signUp(ctx, w, r)

	response.WriteJsonResponse(w, rs)
}

func signUp(ctx context.Context, w http.ResponseWriter, r *http.Request) response.Response {
	log.Println("Sign up")
	empl := repo.Employee{}
	err := json.NewDecoder(r.Body).Decode(&empl)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	_, err = app.DI().EmployeeRepo.CreateEmployee(ctx, empl)
	if err != nil {
		return response.NewServerError("server error")
	}

	return response.NewOkMessageResponse("signed up")
}

func validateEmployee(empl repo.Employee) error {
	if empl.ID == "" {
		return errors.New("invalid id")
	}
	if empl.FirstName == "" {
		return errors.New("invalid first_name")
	}
	if empl.LastName == "" {
		return errors.New("invalid last_name")
	}
	if empl.PositionID == 0 {
		return errors.New("invalid position_id")
	}
	if empl.Login == "" {
		return errors.New("invalid login")
	}
	if empl.Password == "" {
		return errors.New("invalid password")
	}

	return nil
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
