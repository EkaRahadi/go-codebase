package error

import "net/http"

type ForbiddenError struct {
	err error
}

func (e *ForbiddenError) Error() string {
	return e.err.Error()
}

func (e *ForbiddenError) Unwrap() error {
	return e.err
}

func NewForbiddenError(err error) error {
	return NewClientError(&ForbiddenError{err}, http.StatusForbidden)
}
