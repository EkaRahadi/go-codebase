package trace

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type CloseFunc func(ctx context.Context) error

func NewTraceProviderBuilder(name string) *traceProviderBuilder {
	return &traceProviderBuilder{
		name: name,
	}
}

type traceProviderBuilder struct {
	name     string
	exporter trace.SpanExporter
}

func (b *traceProviderBuilder) SetExporter(exp trace.SpanExporter) *traceProviderBuilder {
	b.exporter = exp
	return b
}

func (b *traceProviderBuilder) Build() (*trace.TracerProvider, CloseFunc, error) {
	ctx := context.Background()

	r, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithHost())
	r, err = resource.Merge(resource.Default(), r)
	r, err = resource.Merge(
		r,
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(b.name),
		),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("error build tracer provider")
	}

	if b.exporter == nil {
		return nil, nil, fmt.Errorf("no exporter set to otlp tracer provider")
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(b.exporter),
		trace.WithResource(r),
	)

	return tracerProvider, func(ctx context.Context) error {
		cxt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := tracerProvider.Shutdown(cxt); err != nil {
			return err
		}
		return err
	}, nil
}
