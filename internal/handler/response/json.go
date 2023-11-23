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

func NewOkMessageResponse(message string) Response {
	type ok struct {
		Message string `json:"message"`
	}

	return Response{
		Code:     http.StatusOK,
		ErrorMsg: "",
		Data: ok{
			Message: message,
		},
	}
}

func NewErrResponse(code int, errorMsg string) Response {
	return Response{
		Code:     http.StatusInternalServerError,
		ErrorMsg: errorMsg,
		Data:     nil,
	}
}

func NewServerError(err error) Response {
	log.Printf("Server error: %v\n", err)
	return NewErrResponse(http.StatusInternalServerError, "server error")
}

func NewNotFoundError(errorMsg string) Response {
	return NewErrResponse(http.StatusNotFound, errorMsg)
}

func NewBadRequestErrorResponse(errorMsg string) Response {
	return NewErrResponse(http.StatusBadRequest, errorMsg)
}

func NewUnauthroziedResponse(errorMsg string) Response {
	return NewErrResponse(http.StatusUnauthorized, errorMsg)
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
