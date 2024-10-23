package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/dto"
	apperror "github.com/EkaRahadi/go-codebase/internal/error" //alias
	"github.com/gin-gonic/gin"
)

const (
	MessageInternalServerError        = "currently our server is facing unexpected error, please try again later"
	MessageValidationError            = "input validation error"
	MessageInvalidJsonValueTypeError  = "invalid value for %s"
	MessageInvalidJsonUnmarshallError = "invalid JSON format"
	MessageJsonSyntaxError            = "invalid JSON syntax"
)

// middleware to handle the application errors:
// - client error
// - server error
// unknown errors will be treated as server error, since we don't want to accidentally expose our server error
// while at the same time we want to give our client a clear error message
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 1 {
			//usually not a good idea, but since Gin already throws panic if we give nil to c.Error, it is probably safe
			var err error
			err = c.Errors[0]

			var serverError *apperror.ServerError
			var clientError *apperror.ClientError
			var jutErr *json.UnmarshalTypeError
			var sErr *json.SyntaxError
			var vldtnErr *apperror.ValidationError

			isClientError := false
			if errors.As(err, &clientError) {
				isClientError = true
				err = clientError.Unwrap()
			}

			switch {
			case errors.As(err, &sErr):
				handleJsonSyntaxError(c, sErr)
				return
			case errors.As(err, &jutErr):
				handleJsonUnmarshalTypeError(c, jutErr)
				return
			case isClientError:
				c.AbortWithStatusJSON(clientError.GetCode(), dto.ErrorResponse{
					Message: clientError.Error(),
					Code:    fmt.Sprintf("%d", clientError.GetCode()),
				})
				return
			case errors.As(err, &serverError):
				c.AbortWithStatusJSON(serverError.GetCode(), dto.ErrorResponse{
					Message: MessageInternalServerError,
					Code:    constants.CodeServerError,
				})
				return
			case errors.As(err, &vldtnErr):
				c.AbortWithStatusJSON(vldtnErr.GetCode(), dto.ErrorResponse{
					Message: vldtnErr.Error(),
					Data:    vldtnErr.GetDetails(),
					Code:    constants.CodeBadRequest,
				})
				return

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
					Message: MessageInternalServerError,
					Code:    constants.CodeServerError,
				})
				return
			}
		} else if len(c.Errors) > 1 { // to simplify this case
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Message: MessageInternalServerError,
				Code:    constants.CodeServerError,
			})
			return
		}
	}
}

func handleJsonSyntaxError(c *gin.Context, err *json.SyntaxError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: MessageJsonSyntaxError + strconv.Itoa(int(err.Offset)),
		Code:    constants.CodeBadRequest,
	})
}

func handleJsonUnmarshalTypeError(c *gin.Context, err *json.UnmarshalTypeError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: fmt.Sprintf(MessageInvalidJsonValueTypeError, err.Field),
		Code:    constants.CodeBadRequest,
	})
}
