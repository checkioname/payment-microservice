package application

import (
	"anturiocode/api--order-service/internal/api/protos/api"
	"anturiocode/api--order-service/internal/repository"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type Order struct {
	R *repository.OrderRepository
}

func (o *Order) Save(order *Order) error {
	resp, err := payOrder()
	if err != nil {
		slog.Warn("Erro no processamento gRPC:", err)
		return nil
	}

	if !resp.Success {
		slog.Warn("Erro no processamento do pagamento")
		return errors.New("no processamento do pagamento")
	}

	return nil
}

func payOrder() (*api.PaymentResponse, error) {
	grpcServerAddr := "localhost:50051"

	conn, err := grpc.Dial(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Warn("Erro ao conectar ao servidor gRPC:", err)
		return nil, nil
	}
	defer conn.Close()

	client := api.NewPayerClient(conn)
	req := &api.PaymentRequest{}

	// Chamar o serviço gRPC
	return client.ProcessPayment(context.Background(), req)
}
