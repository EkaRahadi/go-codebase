package error

import "net/http"

type ClientError struct {
	err  error
	code int
}

func (e *ClientError) Error() string {
	return e.err.Error()
}

func (e *ClientError) Unwrap() error {
	return e.err
}

func (e *ClientError) GetCode() int {
	return e.code
}

// wrap error with client error
func NewClientError(err error, code ...int) error {
	statusCode := http.StatusBadRequest

	if len(code) > 0 {
		statusCode = code[0]
	}

	return &ClientError{
		err:  err,
		code: statusCode,
	}
}
