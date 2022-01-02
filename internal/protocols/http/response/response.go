package response

import (
	"encoding/json"
	"golang-starter/internal/protocols/http/errors"
	"net/http"
)

type Response struct {
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func Json(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	res := Response{
		Message: &message,
		Data:    &data,
	}
	json.NewEncoder(w).Encode(res)
}

func Text(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(httpCode)
	w.Write([]byte(message))
}

// TODO: implement response error
func Err(w http.ResponseWriter, err error) {
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	res := Response{
		Message: &er.Message,
	}
	json.NewEncoder(w).Encode(res)
}
