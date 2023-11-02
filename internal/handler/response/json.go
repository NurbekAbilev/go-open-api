package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"error"`
	Data     any    `json:"data"`
}

type ResponseGeneric[T any] struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"error"`
	Data     any    `json:"data"`
}

// type Response struct {
// 	Code     int    `json:"code"`
// 	ErrorMsg string `json:"error"`
// 	Data     any    `json:"data"`
// }

func NewResponse(code int, errorMsg string, data any) Response {
	return Response{
		Code:     code,
		ErrorMsg: errorMsg,
		Data:     data,
	}
}

func NewOkResponse(data any) Response {
	return Response{
		Code:     http.StatusOK,
		ErrorMsg: "",
		Data:     data,
	}
}

func NewErrResponse(code int, errorMsg string) Response {
	return Response{
		Code:     http.StatusInternalServerError,
		ErrorMsg: errorMsg,
		Data:     nil,
	}
}

func NewServerError(errorMsg string) Response {
	return NewErrResponse(http.StatusInternalServerError, errorMsg)
}

func NewBadRequestErrorResponse(errorMsg string) Response {
	return NewErrResponse(http.StatusBadRequest, errorMsg)
}

func WriteJsonResponse(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(r)
	if err != nil {
		log.Printf("error during marshall response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
		return
	}

	w.Write(jsonBytes)
}
