package error

import "net/http"

type TokenError struct {
}

func (e *TokenError) Error() string {
	return "invalid token"
}

func NewTokenError() error {
	return NewClientError(&TokenError{}, http.StatusUnauthorized)
}
