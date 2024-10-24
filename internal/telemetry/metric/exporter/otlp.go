package exporter

import (
	"context"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewOTLP(endpoint string) metric.Exporter {
	ctx := context.Background()
	metricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		// otlpmetricgrpc.WithEndpoint(endpoint),
	)

	if err != nil {
		logger.Log.Fatalw("Failed to create the collector metric exporter", "err", err.Error())
	}
	return metricExporter
}
