package middleware

import (
	"github.com/EkaRahadi/go-codebase/internal/constants"
	apperror "github.com/EkaRahadi/go-codebase/internal/error" //alias
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/gin-gonic/gin"
)

func bindAndValidate[requestStruct any](c *gin.Context, contextName string, binder func(interface{}) error, validator utils.Validator) {
	var body requestStruct

	if err := binder(&body); err != nil {
		c.Error(apperror.NewClientError(err))
		c.Abort()
		return
	}

	if err := validator.Validate(body); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Set(contextName, body)
	c.Next()
}

func JsonBody[requestBodyStruct any](vldtr utils.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		bindAndValidate[requestBodyStruct](c, constants.ContextRequestBodyJSON, c.ShouldBindJSON, vldtr)
	}
}

func Uri[requestUriStruct any](vldtr utils.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		bindAndValidate[requestUriStruct](c, constants.ContextRequestBodyURI, c.ShouldBindUri, vldtr)
	}
}

func Query[requestQueryStruct any](vldtr utils.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		bindAndValidate[requestQueryStruct](c, constants.ContextRequestBodyQuery, c.ShouldBindQuery, vldtr)
	}
}

func Form[requestMultipartStruct any](vldtr utils.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		bindAndValidate[requestMultipartStruct](c, constants.ContextRequestBodyForm, c.ShouldBind, vldtr)
	}
}
