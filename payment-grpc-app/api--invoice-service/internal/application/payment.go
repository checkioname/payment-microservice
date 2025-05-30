package application

import (
	"anturiocode/api--payment-service/internal/api/protos/api"
	"anturiocode/api--payment-service/internal/infrastructure/repositories"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"strconv"
	"time"
)

const (
	max = 1.5
	min = 0.5
)

type InvoiceService struct {
	R repositories.PaymentRepository

	tracer trace.Tracer
}
