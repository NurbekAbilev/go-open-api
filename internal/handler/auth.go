package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nurbekabilev/go-open-api/internal/handler/auth"
)

func HandleAuthEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		io.WriteString(w, "No id provided")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		io.WriteString(w, "No id provided")
		return
	}

	signedToken, err := auth.GenerateSignedJWT(id)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, signedToken)

}

func ValidateAuthEmployee(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	claim := auth.CustomAuthClaims{}

	tkn, err := auth.ValidateSignedJWT(token, &claim)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Token valid:", tkn.Valid)
	log.Println(claim)
}
