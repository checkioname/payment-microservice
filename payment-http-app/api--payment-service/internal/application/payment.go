package application

import (
	"anturiocode/api--payment-service/internal/api/requests"
	"anturiocode/api--payment-service/internal/api/responses"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"time"
)

const (
	max = 1.5
	min = 0.5
)

type PaymentService interface {
	ProcessPayment(context.Context, *requests.PaymentRequest) (*responses.PaymentResponse, error)
}

type paymentService struct {
	tracer trace.Tracer
}

func NewPaymentService(t trace.Tracer) PaymentService {
	return &paymentService{tracer: t}
}

func (i *paymentService) ProcessPayment(context.Context, *requests.PaymentRequest) (*responses.PaymentResponse, error) {
	fmt.Println("Payment called")
	//	# Simula processamento
	n := min + rand.Float64()*(max-min)
	processingTime := time.Duration(n)
	time.Sleep(time.Duration(processingTime) * time.Second)

	select {
	case <-time.After(processingTime):
		success := rand.Float64() > 0.1
		if success {

			return &responses.PaymentResponse{Success: true, Message: "Pagamento com sucesso!"}, nil

		}
		return &responses.PaymentResponse{Success: true, Message: "Pagamento com erro!"}, nil
	}
}
