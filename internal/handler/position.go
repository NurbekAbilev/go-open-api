package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

func HandleAddPosition(w http.ResponseWriter, r *http.Request) {
	di := app.DI()
	ctx := r.Context()

	rs := AddPosition(ctx, di.PositionRepo, r)

	response.WriteJsonResponse(w, rs)
}

func AddPosition(ctx context.Context, createPosRepo repo.CreatePositionRepo, r *http.Request) response.Response {

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

// func (ph PositionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	id, idOk := getIdFromURI(r.RequestURI)

// 	switch r.Method {
// 	case "GET":
// 		// /positions/{id}
// 		if idOk {
// 			handleGetPosition(w, r, ph.DI, id)
// 			return
// 		}

// 		// /positions
// 		handleGetAllPositions(w, r, ph.DI)
// 	case "POST":
// 		handleAddPosition(w, r, ph.DI)
// 	case "PUT":
// 		if idOk {
// 			handleUpdatePosition(w, r, ph.DI, id)
// 		}
// 	default:
// 		fmt.Fprint(w, "Invalid method")
// 	}
// }

// func getIdFromURI(uri string) (id int, ok bool) {
// 	uriParts := strings.Split(uri, "/")
// 	if len(uriParts) < 4 {
// 		return 0, false
// 	}
// 	idPart := uriParts[4]

// 	id, err := strconv.Atoi(idPart)
// 	if err != nil {
// 		return 0, false
// 	}

// 	return id, true
// }

// func handleGetPosition(w http.ResponseWriter, r *http.Request, di Injector, id int) {
// 	repo := repo.PositionRepo{}
// 	pos, err := repo.GetPositionById(di.DB, id)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	resp.WriteJsonResponse(w, false, "", pos)
// }

// func handleGetAllPositions(w http.ResponseWriter, r *http.Request, di Injector) {
// 	fmt.Fprintf(w, "handleGetPosition")
// }

// func handleUpdatePosition(w http.ResponseWriter, r *http.Request, di Injector, id int) {
// 	fmt.Fprintf(w, "handleGetPosition ", id)
// }
