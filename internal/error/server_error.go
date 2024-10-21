package error

import "net/http"

type ServerError struct {
	err  error
	code int
}

func (e *ServerError) Error() string {
	return e.err.Error()
}

func (e *ServerError) Unwrap() error {
	return e.err
}

func (e *ServerError) GetCode() int {
	return e.code
}

func NewServerError(err error, code ...int) error {
	statusCode := http.StatusInternalServerError

	if len(code) > 0 {
		statusCode = code[0]
	}

	return &ServerError{
		err:  err,
		code: statusCode,
	}
}
