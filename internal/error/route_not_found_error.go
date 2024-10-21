package error

import "net/http"

type RouteNotFoundError struct {
}

func (e RouteNotFoundError) Error() string {
	return "route not found"
}

// wrapping route not found error with client error
func NewRouteNotFoundError() error {
	return NewClientError(RouteNotFoundError{}, http.StatusNotFound)
}
