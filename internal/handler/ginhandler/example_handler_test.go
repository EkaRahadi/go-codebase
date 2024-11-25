package ginhandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	"github.com/EkaRahadi/go-codebase/internal/handler/ginhandler"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/EkaRahadi/go-codebase/internal/utils/testutil"
	mockHttpClient "github.com/EkaRahadi/go-codebase/mocks/httpclient"
	mockUsecase "github.com/EkaRahadi/go-codebase/mocks/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestExampleHandlerFunc(t *testing.T) {
	t.Run("should return success with code 200", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := testutil.SetupRouter()
		validator := utils.NewCustomValidator()
		mockExampleUC := mockUsecase.NewExampleUsecase(t)
		dummy := &entity.Dummy{
			Message: "Test Dummy",
		}
		mockExampleUC.On("ExampleUCFunc", mock.Anything).Return(dummy, nil)
		mockHtppClient := mockHttpClient.NewHttpClient(t)
		testHttpResCall := "test response body"
		mockHtppClient.On("GetWithQuery", mock.Anything, mock.Anything, mock.Anything).Return(testHttpResCall, nil)
		handler := ginhandler.NewExampleHandler(mockExampleUC, mockHtppClient)
		router.POST("/example", middleware.JsonBody[dto.DummyRequest](validator), handler.ExampleHandlerFunc)
		reqBody := &dto.DummyRequest{
			Foo: "Hello",
			Bar: "World",
		}
		request, _ := json.Marshal(reqBody)
		expected, _ := json.Marshal(dto.SuccessResponse{
			Message: "success",
			Data: map[string]string{
				"jsonplaceholder": testHttpResCall,
				"message":         dummy.Message,
			},
			Code: "200",
		})

		req, _ := http.NewRequest(http.MethodPost, "/example", bytes.NewBuffer(request))
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	})
}
