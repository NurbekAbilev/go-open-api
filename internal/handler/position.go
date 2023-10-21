package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type PositionHandler struct {
	DI Injector
}

func (ph PositionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Position handler:")
	if r.Method == "POST" {
		handleAddPosition(w, r, ph.DI)
		return
	}
	if r.Method == "GET" {
		handleGetOnePosition(w, r, ph.DI)
		return
	}

	fmt.Fprint(w, "Invalid method")
}

func validateAddPositionStruct(p repo.Position) error {
	if p.Name == nil || *p.Name == "" {
		return errors.New("invalid name")
	}

	if p.Salary == nil || *p.Salary < 0 {
		return errors.New("invalid salary")
	}

	return nil
}

func handleAddPosition(w http.ResponseWriter, r *http.Request, di Injector) {
	p := repo.Position{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintf(w, "could not decode json body")
		return
	}

	err = validateAddPositionStruct(p)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	repo := repo.PositionRepo{}
	err = repo.CreatePosition(di.DB, p)
	if err != nil {
		fmt.Fprintf(w, "could not create position")
		return
	}

	fmt.Fprintf(w, "record created succesfully")
}

func handleGetOnePosition(w http.ResponseWriter, r *http.Request, di Injector) {
}
