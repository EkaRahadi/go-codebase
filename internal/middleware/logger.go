package middleware

import (
	"time"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/EkaRahadi/go-codebase/internal/telemetry"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func recordRequest(ctx *gin.Context, counter metric.Int64Counter) {
	// Add the label to the metrics
	labels := []metric.AddOption{
		metric.WithAttributes(semconv.HTTPResponseStatusCode(ctx.Writer.Status())),
		metric.WithAttributes(semconv.HTTPMethod(ctx.Request.Method)),
		metric.WithAttributes(semconv.URLPath(ctx.Request.URL.Path)),
	}

	counter.Add(ctx, 1, labels...)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		reqCount, err := telemetry.Meter.Int64Counter("http.server.requests",
			metric.WithDescription("Count how many request processed"),
			metric.WithUnit("1"))
		if err != nil {
			logger.Log.Errorw("failed to create metric", "meta", err.Error())
		}

		c.Next()

		param := map[string]interface{}{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency_ms":  time.Since(start).Milliseconds(),
			"path":        path,
		}

		// Metric Instrumentation
		recordRequest(c, reqCount)

		if len(c.Errors) == 0 {
			logger.Log.Infow("request success", "meta", param)
		} else {
			errList := []error{}
			for _, err := range c.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				param["errors"] = errList
				logger.Log.Errorw("request error", "meta", param)
			}
		}
	}
}
