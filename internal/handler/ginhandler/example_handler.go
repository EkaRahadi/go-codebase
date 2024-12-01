package ginhandler

import (
	"log"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/helper/request"
	"github.com/EkaRahadi/go-codebase/internal/helper/response"
	"github.com/EkaRahadi/go-codebase/internal/httpclient"
	"github.com/EkaRahadi/go-codebase/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleUsecase usecase.ExampleUsecase
	httpclient     httpclient.HttpClient
}

func NewExampleHandler(exampleUsecase usecase.ExampleUsecase, httpclient httpclient.HttpClient) *ExampleHandler {
	return &ExampleHandler{
		exampleUsecase: exampleUsecase,
		httpclient:     httpclient,
	}
}

func (h *ExampleHandler) ExampleHandlerFunc(c *gin.Context) {
	_ = request.GetJsonRequestBody[dto.DummyRequest](c)
	ctx := c.Request.Context() // make sure Extract parent context to enabled distributed tracing whenever use httpclient
	resGet, err := h.httpclient.GetWithQuery(ctx, "http://localhost:8080/example-with-tx", map[string]string{
		"foo": "1",
		"bar": "2",
	})
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.exampleUsecase.ExampleUCFunc(c)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, map[string]string{
		"jsonplaceholder": resGet,
		"message":         res.Message,
	})
}

func (h *ExampleHandler) ExampleHandlerWithTxFunc(c *gin.Context) {
	requestQuery := request.GetQueryRequest[dto.DummyRequestQuery](c)
	log.Println("requestQuery", requestQuery)

	res, err := h.exampleUsecase.ExampleUCTXFunc(c)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, res)
}

func (h *ExampleHandler) ExampleHandlerFuncUri(c *gin.Context) {
	requestUri := request.GetUriRequest[dto.DummyRequestUri](c)
	log.Println("requestUri", requestUri)

	res, err := h.exampleUsecase.ExampleUCFunc(c)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, res)
}
