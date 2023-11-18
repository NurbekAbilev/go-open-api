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
	user := repo.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	err = validateUser(user)
	if err != nil {
		return response.NewBadRequestErrorResponse(err.Error())
	}

	const cost = 8
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return response.NewServerError("server error")
	}

	user.Password = string(hash)

	_, err = app.DI().UserRepo.CreateUser(ctx, user)
	if err != nil {
		return response.NewServerError("server error")
	}

	return response.NewOkMessageResponse("signed up")
}

func validateUser(user repo.User) error {
	if user.Login == "" {
		return errors.New("invalid last_name")
	}
	if user.Password == "" {
		return errors.New("invalid position_id")
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

	user, err := app.DI().UserRepo.GetUserByLogin(ctx, rb.Login)
	if err != nil {
		return response.NewBadRequestErrorResponse("No employee found by login/password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rb.Password))
	if err != nil {
		return response.NewBadRequestErrorResponse("No employee found by login/password")
	}

	cred := auth.Credentials{
		Login: user.Login,
		ID:    user.ID,
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
