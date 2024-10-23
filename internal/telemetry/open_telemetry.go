package telemetry

import (
	"context"

	"github.com/EkaRahadi/go-codebase/internal/logger"
	"github.com/EkaRahadi/go-codebase/internal/telemetry/metric"
	metricExporter "github.com/EkaRahadi/go-codebase/internal/telemetry/metric/exporter"
	ttrace "github.com/EkaRahadi/go-codebase/internal/telemetry/trace"
	traceExporter "github.com/EkaRahadi/go-codebase/internal/telemetry/trace/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (t *Telemetry) InitGlobalProviderOpenTelemetry(endpoint string, appName string) {
	metricExp := metricExporter.NewOTLP(endpoint)
	meterProvider, metricCloseFn, err := metric.NewMeterProviderBuilder().
		SetExporter(metricExp).
		Build(appName)
	if err != nil {
		logger.Log.Fatalw("failed initializing the meter provider", "error", err)
	}

	t.metricProviderCloseFn = append(t.metricProviderCloseFn, metricCloseFn)

	spanExporter := traceExporter.NewOTLP(endpoint)
	tracerProvider, tracerProviderCloseFn, err := ttrace.NewTraceProviderBuilder(appName).
		SetExporter(spanExporter).
		Build()
	if err != nil {
		logger.Log.Fatalw("failed initializing the tracer provider", "error", err)
	}
	t.traceProviderCloseFn = append(t.traceProviderCloseFn, tracerProviderCloseFn)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetMeterProvider(meterProvider)
	otel.SetTracerProvider(tracerProvider)
}

func (t Telemetry) Shutdown(ctxShutDown context.Context) {
	for _, closeFn := range t.metricProviderCloseFn {
		go func() {
			err := closeFn(ctxShutDown)
			if err != nil {
				logger.Log.Errorw("Unable to close metric provider")
			}
		}()
	}
	for _, closeFn := range t.traceProviderCloseFn {
		go func() {
			err := closeFn(ctxShutDown)
			if err != nil {
				logger.Log.Errorw("Unable to close trace provider")
			}
		}()
	}
}
