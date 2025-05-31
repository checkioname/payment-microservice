package client

import (
	order2 "anturiocode/api--order-service/internal/api/protos/order"
	"anturiocode/api--order-service/internal/api/protos/payment"
	payment2 "anturiocode/api--order-service/internal/api/protos/payment"
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	client payment.PayerClient
	conn   *grpc.ClientConn
}

func NewPaymentClient(addr string) *PaymentClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Erro ao criar client do inventory", err)
		return nil
	}
	client := payment.NewPayerClient(conn)
	return &PaymentClient{client: client, conn: conn}
}

func (oc *PaymentClient) PayOrder(req *order2.CreateOrderRequest) (*payment2.PaymentResponse, error) {
	grpcServerAddr := ":8008"

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
