package testutil

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.ErrorHandler(),
		gin.Recovery(),
	}
	router.Use(middlewares...)
	return router
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
