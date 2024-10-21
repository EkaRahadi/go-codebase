package error

import (
	"net/http"
)

type UnauthorizedError struct {
	err error
}

func (e *UnauthorizedError) Error() string {
	return e.err.Error()
}

func (e *UnauthorizedError) Unwrap() error {
	return e.err
}

func NewUnauthorizedError(err error) error {
	return NewClientError(&UnauthorizedError{err}, http.StatusUnauthorized)
}
