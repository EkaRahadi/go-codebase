package dto

type ValidationErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}
