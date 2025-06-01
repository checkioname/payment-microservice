package application

import (
	inventory "anturiocode/api--inventory-service/internal/api/protos"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"

	"time"
)

type InventoryService struct {
	tracer trace.Tracer
	inventory.UnimplementedInventoryServiceServer
}

func NewInventoryService(t trace.Tracer) inventory.InventoryServiceServer {
	return &InventoryService{tracer: t}
}

func (i InventoryService) CheckAndReserveStock(context.Context, *inventory.CheckAndReserveStockRequest) (*inventory.CheckAndReserveStockResponse, error) {
	fmt.Println("Dispatch called")
	time.Sleep(1 * time.Second)

	return &inventory.CheckAndReserveStockResponse{Success: true, Message: "Produto verificado no estoque"}, nil
}

func (i InventoryService) ReleaseStock(context.Context, *inventory.ReleaseStockRequest) (*inventory.ReleaseStockResponse, error) {
	fmt.Println("Release called")
	time.Sleep(1 * time.Second)
	return &inventory.ReleaseStockResponse{Success: true, Message: "Produto reservado"}, nil
}
