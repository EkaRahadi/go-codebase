package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/EkaRahadi/go-codebase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogger_SuccessfulRequest(t *testing.T) {
	mockLogger := mocks.NewLogger(t)
	logger.Log = mockLogger
	mockLogger.On("Infow", "incoming request success", "meta", mock.Anything).Run(func(args mock.Arguments) {
		meta := args.Get(2).(map[string]interface{})
		assert.Equal(t, http.StatusOK, meta["status_code"])
		assert.Equal(t, "GET", meta["method"])
		assert.Equal(t, "/test", meta["path"])
		assert.GreaterOrEqual(t, meta["latency_ms"].(int64), int64(0))
	}).Once()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.Logger())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockLogger.AssertExpectations(t) // Assert that the logger was called
}

func TestLogger_RequestWithError(t *testing.T) {
	mockLogger := mocks.NewLogger(t)
	logger.Log = mockLogger
	mockLogger.On("Errorw", "incoming request having error", "meta", mock.Anything).Run(func(args mock.Arguments) {
		meta := args.Get(2).(map[string]interface{})
		assert.Equal(t, http.StatusInternalServerError, meta["status_code"])
		assert.Equal(t, "GET", meta["method"])
		assert.Equal(t, "/test", meta["path"])
		assert.GreaterOrEqual(t, meta["latency_ms"].(int64), int64(0))
		assert.NotEmpty(t, meta["errors"])
	}).Once()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.Logger())
	router.GET("/test", func(c *gin.Context) {
		c.Error(assert.AnError) // Simulate an error
		c.AbortWithStatus(http.StatusInternalServerError)
	})

	// Perform the request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Assert that the logger was called
	mockLogger.AssertExpectations(t)
}
