package telemetry

import (
	"github.com/EkaRahadi/go-codebase/internal/telemetry/metric"
	ttrace "github.com/EkaRahadi/go-codebase/internal/telemetry/trace"
)

type Telemetry struct {
	metricProviderCloseFn []metric.CloseFunc
	traceProviderCloseFn  []ttrace.CloseFunc
}

func NewTelemetry() *Telemetry {
	return &Telemetry{}
}
