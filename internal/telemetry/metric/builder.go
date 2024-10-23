package metric

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type CloseFunc func(ctx context.Context) error

func NewMeterProviderBuilder() *meterProviderBuilder {
	return &meterProviderBuilder{}
}

type meterProviderBuilder struct {
	exporter metric.Exporter
}

func (b *meterProviderBuilder) SetExporter(exp metric.Exporter) *meterProviderBuilder {
	b.exporter = exp
	return b
}

func (b meterProviderBuilder) Build(serviceName string) (*metric.MeterProvider, CloseFunc, error) {
	res, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			// semconv.ServiceVersion("0.1.0"),
		))
	if err != nil {
		return nil, nil, fmt.Errorf("error build meter provider")
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(b.exporter,
			// Default is 1m.
			metric.WithInterval(15*time.Second))),
	)

	if b.exporter == nil {
		return nil, nil, fmt.Errorf("no exporter set to otlp metric provider")
	}

	return meterProvider, func(ctx context.Context) error {
		cxt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := meterProvider.Shutdown(cxt); err != nil {
			return err
		}
		return nil
	}, nil

}
