package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/auth"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
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
	empl := repo.Employee{}
	err := json.NewDecoder(r.Body).Decode(&empl)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	err = validateEmployee(empl)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	const cost = 8
	hash, err := bcrypt.GenerateFromPassword([]byte(empl.Password), cost)
	if err != nil {
		return response.NewServerError("server error")
	}

	empl.Password = string(hash)

	_, err = app.DI().EmployeeRepo.CreateEmployee(ctx, empl)
	if err != nil {
		return response.NewServerError("server error")
	}

	return response.NewOkMessageResponse("signed up")
}

func validateEmployee(empl repo.Employee) error {
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
	type requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	rb := requestBody{}

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		return response.NewBadRequestErrorResponse("Invalid json %w")
	}

	empl, err := app.DI().EmployeeRepo.FindEmployeeByLogin(ctx, rb.Login)
	if err != nil {
		return response.NewBadRequestErrorResponse("No employee found by login/password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(empl.Password), []byte(rb.Password))
	if err != nil {
		return response.NewBadRequestErrorResponse("No employee found by login/password")
	}

	cred := auth.Credentials{
		Login: empl.Login,
		ID:    empl.ID,
	}

	token, err := authProvider.GenerateToken(cred)
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
