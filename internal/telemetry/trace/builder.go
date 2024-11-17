package trace

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/helper/service"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
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
			semconv.ServiceInstanceID(service.GetInstanceID()),
			attribute.String("k8s.namespace", os.Getenv("POD_NAMESPACE")),
			semconv.K8SPodNameKey.String(os.Getenv("HOSTNAME")), //podname
		),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("error build tracer provider: %w", err)
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
