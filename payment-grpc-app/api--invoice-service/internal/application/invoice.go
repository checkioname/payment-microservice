package application

import (
	"go.opentelemetry.io/otel/trace"
)

const (
	max = 1.5
	min = 0.5
)

type InvoiceService struct {
	tracer trace.Tracer
}
