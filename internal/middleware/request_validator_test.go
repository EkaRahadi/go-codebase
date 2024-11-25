package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	mockUtils "github.com/EkaRahadi/go-codebase/mocks/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RequestBody struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func TestJsonBodyMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockValidator := mockUtils.NewValidator(t)
	router := gin.New()
	router.Use(middleware.JsonBody[RequestBody](mockValidator))
	router.POST("/test", func(c *gin.Context) {
		data, exists := c.Get(constants.ContextRequestBodyJSON)
		assert.True(t, exists, "Request body should be set in context")
		requestBody, ok := data.(RequestBody)
		assert.True(t, ok, "Data in context should be of type RequestBody")
		assert.Equal(t, "John Doe", requestBody.Name)
		assert.Equal(t, 30, requestBody.Age)

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	validInput := `{"name": "John Doe", "age": 30}`
	mockValidator.On("Validate", mock.Anything).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(validInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockValidator.AssertCalled(t, "Validate", RequestBody{Name: "John Doe", Age: 30})
}
