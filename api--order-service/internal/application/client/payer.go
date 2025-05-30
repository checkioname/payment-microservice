package client

import (
	order2 "anturiocode/api--order-service/internal/api/protos/order"
	payment2 "anturiocode/api--order-service/internal/api/protos/payment"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type OrderClient struct {
}

func (oc *OrderClient) PayOrder(req *order2.CreateOrderRequest) (*payment2.PaymentResponse, error) {
	grpcServerAddr := "localhost:50051"

	conn, err := grpc.NewClient(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Warn("Erro ao conectar ao servidor gRPC:", err)
		return nil, nil
	}
	defer conn.Close()

	client := payment2.NewPayerClient(conn)
	r := &payment2.PaymentRequest{
		OrderId: req.CustomerId,
	}
	return client.ProcessPayment(context.Background(), r)
}
