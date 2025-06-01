package application

import (
	"anturiocode/api--invoice-service/internal/api/requests"
	"anturiocode/api--invoice-service/internal/api/responses"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"math/rand"

	"time"
)

type InvoiceService interface {
	GetInvoice(context.Context, *requests.InvoiceRequest) (*responses.InvoiceResponse, error)
}

type invoiceService struct {
	tracer trace.Tracer
}

func NewInvoiceService(t trace.Tracer) InvoiceService {
	return &invoiceService{tracer: t}
}

func (i *invoiceService) GetInvoice(context.Context, *requests.InvoiceRequest) (*responses.InvoiceResponse, error) {
	fmt.Println("GetInvoice called")
	time.Sleep(1 * time.Second)
	return &responses.InvoiceResponse{Success: true, Message: fmt.Sprintf("Numero da Nota: d%", rand.Int())}, nil

}
