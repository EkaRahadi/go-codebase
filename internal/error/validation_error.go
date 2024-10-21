package error

import (
	"net/http"

	"github.com/EkaRahadi/go-codebase/internal/dto"
)

type ValidationError struct {
	message string
	details []dto.ValidationErrorResponse
	code    int
}

func (e *ValidationError) Error() string {
	return e.message
}

func (e *ValidationError) GetCode() int {
	return e.code
}

func (e *ValidationError) GetDetails() []dto.ValidationErrorResponse {
	return e.details
}

func NewValidationError(message string, details []dto.ValidationErrorResponse, code ...int) error {
	statusCode := http.StatusBadRequest

	if len(code) > 0 {
		statusCode = code[0]
	}

	return &ValidationError{
		message: message,
		details: details,
		code:    statusCode,
	}
}
