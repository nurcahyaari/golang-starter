package errors

import (
	"fmt"
	"net/http"
	"regexp"
)

type RespError struct {
	Code    int
	Message string
}

func (r *RespError) Error() string {
	return fmt.Sprintf("%d: %d", r.Code, r.Message)
}

func InternalServerError(msg string) error {
	return &RespError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func BadRequest(msg string) error {
	return &RespError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NotFound(msg string) error {
	return &RespError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func Unauthorization(msg string) error {
	return &RespError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func FindErrorType(err error) error {
	re := regexp.MustCompile(`not found.?`)
	if re.FindString(err.Error()) != "" {
		return NotFound(err.Error())
	}

	return InternalServerError(err.Error())
}
