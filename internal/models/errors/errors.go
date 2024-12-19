package errors

import (
	"fmt"
	"net/http"
)

func ErrInvalidBodyJSON(err error) HttpAppError {
	return NewHttpAppError(http.StatusBadRequest, err.Error())
}

func ErrUnexpected(err error) HttpAppError {
	return NewHttpAppError(http.StatusInternalServerError, err.Error())
}

func ErrComponentNotFound(component string) HttpAppError {
	return NewHttpAppError(http.StatusNotFound, fmt.Sprintf("%s not found", component))
}
