package application

import (
	"anturiocode/api--payment-service/internal/api/protos/api"
	"anturiocode/api--payment-service/internal/infrastructure/observability"
	"anturiocode/api--payment-service/internal/infrastructure/repositories"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	max = 1.5
	min = 0.5
)

var paymentRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "payment_requests_total",
		Help: "Total de chamadas ao método ProcessPayment",
	},
)

type PaymentService struct {
	R       repositories.PaymentRepository
	tracer  trace.Tracer
	metrics *observability.Metrics
	api.UnimplementedPayerServer
}

func NewPaymentService(r repositories.PaymentRepository, t trace.Tracer, m *observability.Metrics) api.PayerServer {
	return &PaymentService{R: r, tracer: t, metrics: m}
}

func (p PaymentService) ProcessPayment(ctx context.Context, request *api.PaymentRequest) (*api.PaymentResponse, error) {
	start := time.Now()

	// gera um id para rastrear a request
	ctx, span := p.tracer.Start(ctx, "Service: ProcessPayment",
		trace.WithAttributes(attribute.String("request.id", strconv.Itoa(int(request.OrderId)))),
	)
	defer span.End()

	//	# Simula processamento
	n := min + rand.Float64()*(max-min)
	processingTime := time.Duration(n)
	time.Sleep(time.Duration(processingTime) * time.Second)

	select {
	case <-time.After(processingTime):
		success := rand.Float64() > 0.1
		if success {
			paymentRequests.Inc()
			span.SetAttributes(attribute.Bool("payment.success", true))
			p.R.RegisterPayment(10, 1, ctx)

			duration := time.Since(start).Seconds()
			p.metrics.ObserveDuration("200", "ProcessPayment", duration)
			return &api.PaymentResponse{Success: true, Message: "Tudo certo com o pagamento"}, nil
		}
		span.SetAttributes(attribute.Bool("payment.success", false))
		span.RecordError(fmt.Errorf("payment processing failed"))
		return &api.PaymentResponse{Success: false, Message: "Houve um erro no pagamento"}, fmt.Errorf("payment processing failed")

	case <-ctx.Done():
		span.SetStatus(1, "Context canceled")
		return nil, ctx.Err()
	}
}
