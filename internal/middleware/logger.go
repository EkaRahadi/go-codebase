package middleware

import (
	"time"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		param := map[string]interface{}{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
		}

		if len(c.Errors) == 0 {
			logger.Log.Infow("incoming request success", "meta", param)
		} else {
			errList := []error{}
			for _, err := range c.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				param["errors"] = errList
				logger.Log.Errorw("incoming request having error", "meta", param)
			}
		}
	}
}
