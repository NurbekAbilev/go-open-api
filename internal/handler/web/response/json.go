package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseData struct {
	Err      bool        `json:"error"`
	Messaage string      `json:"error_message"`
	Data     interface{} `json:"data"`
}

func WriteJsonResponse(w http.ResponseWriter, isError bool, message string, data interface{}) error {
	responseApiObj := responseData{
		Err:      isError,
		Messaage: message,
		Data:     data,
	}

	response, err := json.Marshal(responseApiObj)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		return err
	}

	return nil
}
