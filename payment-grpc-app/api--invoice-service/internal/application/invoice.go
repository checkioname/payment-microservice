package application

import (
	invoice "anturiocode/api--invoice-service/internal/api/protos"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"time"
)

type InvoiceService struct {
	tracer trace.Tracer
	invoice.UnimplementedInvoicerServer
}

func NewInvoiceService(t trace.Tracer) invoice.InvoicerServer {
	return &InvoiceService{tracer: t}
}

func (i InvoiceService) GetInvoice(context.Context, *invoice.InvoiceRequest) (*invoice.InvoiceResponse, error) {
	fmt.Println("GetInvoice called")
	time.Sleep(1 * time.Second)
	return &invoice.InvoiceResponse{Success: true, Message: fmt.Sprintf("Numero da Nota: d%", rand.Int())}, nil
}
