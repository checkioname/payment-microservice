package application

import (
	"anturiocode/api--order-service/infrastructure/repositories"
	"anturiocode/api--order-service/internal/api/protos/api"
	"context"
	"errors"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type OrderService struct {
	R repositories.OrderRepository
	api.UnimplementedOrderServiceServer
	t trace.Tracer
}

func NewOrderService(r repositories.OrderRepository, tracer trace.Tracer) api.OrderServiceServer {
	return &OrderService{R: r, t: tracer}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	resp, err := payOrder(req)
	if err != nil {
		slog.Warn("Erro no processamento gRPC:", err)
		return &api.CreateOrderResponse{}, errors.New("fail to create order")
	}

	if !resp.Success {
		slog.Warn("Erro no processamento do pagamento")
		return &api.CreateOrderResponse{}, errors.New("no processamento do pagamento")
	}

	// Registra um pedido no banco e parte para o estoque

	// Separa o produto no estoque (chama API de estoque) e chama API de nota fiscal numa go routine

	// Chama API de logistica para iniciar entrega (devolver codigo de rastreio

	return &api.CreateOrderResponse{}, nil
}

func (o *OrderService) UpdateOrderStatus(context.Context, *api.UpdateOrderStatusRequest) (*api.UpdateOrderStatusResponse, error) {
	return &api.UpdateOrderStatusResponse{}, nil
}

func payOrder(req *api.CreateOrderRequest) (*api.PaymentResponse, error) {
	grpcServerAddr := "localhost:50051"

	conn, err := grpc.NewClient(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Warn("Erro ao conectar ao servidor gRPC:", err)
		return nil, nil
	}
	defer conn.Close()

	client := api.NewPayerClient(conn)
	r := &api.PaymentRequest{
		OrderId: req.CustomerId,
	}
	return client.ProcessPayment(context.Background(), r)
}
