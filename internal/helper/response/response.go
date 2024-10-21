package response

import (
	"fmt"
	"net/http"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/gin-gonic/gin"
)

func ResponseOKPlain(c *gin.Context) {
	ResponseOKData(c, nil)
}

func ResponseOKData(c *gin.Context, data interface{}) {
	ResponseOK(c, "success", data)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	ResponseSuccessJSON(c, http.StatusOK, message, data)
}

func ResponseSuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse{
		Message: message,
		Data:    data,
		Code:    fmt.Sprintf("%d", statusCode),
	})
}

func ResponseSuccessJSONCustom(c *gin.Context, statusCode int, message string, code string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse{
		Message: message,
		Data:    data,
		Code:    code,
	})
}
