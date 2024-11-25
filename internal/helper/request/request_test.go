package request_test

import (
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/helper/request"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type MockRequestBody struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestGetJsonRequestBody(t *testing.T) {
	// Create a mock Gin context
	c := &gin.Context{}
	c.Set(constants.ContextRequestBodyJSON, MockRequestBody{
		Field1: "value1",
		Field2: 123,
	})

	// Call the function
	result := request.GetJsonRequestBody[MockRequestBody](c)

	// Assert the result
	assert.Equal(t, "value1", result.Field1)
	assert.Equal(t, 123, result.Field2)
}

func TestGetUriRequest(t *testing.T) {
	// Create a mock Gin context
	c := &gin.Context{}
	c.Set(constants.ContextRequestBodyURI, MockRequestBody{
		Field1: "uriValue",
		Field2: 456,
	})

	// Call the function
	result := request.GetUriRequest[MockRequestBody](c)

	// Assert the result
	assert.Equal(t, "uriValue", result.Field1)
	assert.Equal(t, 456, result.Field2)
}

func TestGetQueryRequest(t *testing.T) {
	// Create a mock Gin context
	c := &gin.Context{}
	c.Set(constants.ContextRequestBodyQuery, MockRequestBody{
		Field1: "queryValue",
		Field2: 789,
	})

	// Call the function
	result := request.GetQueryRequest[MockRequestBody](c)

	// Assert the result
	assert.Equal(t, "queryValue", result.Field1)
	assert.Equal(t, 789, result.Field2)
}

func TestGetFormRequest(t *testing.T) {
	// Create a mock Gin context
	c := &gin.Context{}
	c.Set(constants.ContextRequestBodyForm, MockRequestBody{
		Field1: "formValue",
		Field2: 1011,
	})

	// Call the function
	result := request.GetFormRequest[MockRequestBody](c)

	// Assert the result
	assert.Equal(t, "formValue", result.Field1)
	assert.Equal(t, 1011, result.Field2)
}
