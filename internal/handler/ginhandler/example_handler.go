package ginhandler

import (
	"log"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/helper/request"
	"github.com/EkaRahadi/go-codebase/internal/helper/response"
	"github.com/EkaRahadi/go-codebase/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleUsecase usecase.ExampleUsecase
}

func NewExampleHandler(exampleUsecase usecase.ExampleUsecase) *ExampleHandler {

	return &ExampleHandler{
		exampleUsecase: exampleUsecase,
	}
}

func (h *ExampleHandler) ExampleHandlerFunc(c *gin.Context) {
	requestBody := request.GetJsonRequestBody[dto.DummyRequest](c)
	log.Println("requestBody", requestBody)

	res, err := h.exampleUsecase.ExampleUCFunc(c)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, res)
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
