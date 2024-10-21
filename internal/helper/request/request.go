package request

import (
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/gin-gonic/gin"
)

func GetJsonRequestBody[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestBodyJSON).(bodyRequestType)
}

func GetUriRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestBodyURI).(bodyRequestType)
}

func GetQueryRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestBodyQuery).(bodyRequestType)
}

func GetFormRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestBodyForm).(bodyRequestType)
}
