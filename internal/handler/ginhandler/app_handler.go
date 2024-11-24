package ginhandler

import (
	apperror "github.com/EkaRahadi/go-codebase/internal/error"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) RouteNotFound(c *gin.Context) {
	//nolint:errcheck // no need to check
	c.Error(apperror.NewRouteNotFoundError())
}
