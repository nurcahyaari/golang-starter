package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error   *string      `json:"error,omitempty"`
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

}
