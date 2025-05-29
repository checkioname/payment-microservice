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

type PaymentService struct {
	R repositories.PaymentRepository
	api.UnimplementedPayerServer
	tracer trace.Tracer
}

func NewPaymentService(r repositories.PaymentRepository, t trace.Tracer) api.PayerServer {
	return &PaymentService{R: r, tracer: t}
}

func (p PaymentService) ProcessPayment(ctx context.Context, request *api.PaymentRequest) (*api.PaymentResponse, error) {
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
			span.SetAttributes(attribute.Bool("payment.success", true))
			p.R.RegisterPayment(10, 1, ctx)
			return &api.PaymentResponse{}, nil
		}
		span.SetAttributes(attribute.Bool("payment.success", false))
		span.RecordError(fmt.Errorf("payment processing failed"))
		return &api.PaymentResponse{}, fmt.Errorf("payment processing failed")

	case <-ctx.Done():
		span.SetStatus(1, "Context canceled")
		return nil, ctx.Err()
	}
}
