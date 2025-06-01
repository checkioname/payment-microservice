package application

import (
	"anturiocode/api--inventory-service/internal/api/requests"
	"anturiocode/api--inventory-service/internal/api/responses"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"

	"time"
)

type InventoryService interface {
	CheckAndReserveStock(context.Context, *requests.CheckAndReserveStockRequest) (*responses.CheckAndReserveStockResponse, error)
}

type inventoryService struct {
	tracer trace.Tracer
}

func NewInventoryService(t trace.Tracer) InventoryService {
	return &inventoryService{tracer: t}
}

func (i *inventoryService) CheckAndReserveStock(context.Context, *requests.CheckAndReserveStockRequest) (*responses.CheckAndReserveStockResponse, error) {
	fmt.Println("Dispatch called")
	time.Sleep(1 * time.Second)

	return &responses.CheckAndReserveStockResponse{Success: true, Message: "Produto verificado no estoque"}, nil
}
