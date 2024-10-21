package middleware

import (
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// adding request-id using uuid to context
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New().String()
		c.Set(constants.ContextKeyRequestId, uuid)
		c.Next()
	}
}
