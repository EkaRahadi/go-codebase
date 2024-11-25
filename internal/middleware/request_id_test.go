package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestIdMiddlewareUsingGin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.RequestId())
	router.GET("/test", func(c *gin.Context) {
		requestId, exists := c.Get(constants.ContextKeyRequestId)

		assert.True(t, exists, "Request ID should exist in the context")
		assert.NotEmpty(t, requestId, "Request ID should not be empty")

		c.JSON(http.StatusOK, gin.H{"request_id": requestId})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "HTTP status should be 200")
	response := make(map[string]string)
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.NotEmpty(t, response["request_id"], "Response should include a request ID")
}
