package telemetry

import (
	"go.opentelemetry.io/otel"
)

var Tracer = otel.Tracer("")
var Meter = otel.Meter("")
