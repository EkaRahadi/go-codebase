package telemetry

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("")
var meter = otel.Meter("")
