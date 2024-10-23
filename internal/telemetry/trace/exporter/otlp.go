package exporter

import (
	"context"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewOTLP(endpoint string) trace.SpanExporter {
	ctx := context.Background()
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint))
	if err != nil {
		logger.Log.Fatalw("Failed to create the collector trace exporter", "error", err)
	}
	return traceExporter
}
