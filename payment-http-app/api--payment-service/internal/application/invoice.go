package application

import (
	"anturiocode/api--payment-service/internal/api/requests"
	"anturiocode/api--payment-service/internal/api/responses"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type PaymentService interface {
	HandlePayment(context.Context, *requests.PaymentRequest) (*responses.PaymentResponse, error)
}

type paymentService struct {
	tracer trace.Tracer
}

func NewInventoryService(t trace.Tracer) PaymentService {
	return &paymentService{tracer: t}
}

func (i *paymentService) HandlePayment(context.Context, *requests.PaymentRequest) (*responses.PaymentResponse, error) {
	fmt.Println("GetInvoice called")
	time.Sleep(1 * time.Second)
	return &responses.PaymentResponse{Success: true, Message: "Pagamento com sucesso!"}, nil
}
