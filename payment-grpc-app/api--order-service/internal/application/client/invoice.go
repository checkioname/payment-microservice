package client

import (
	"anturiocode/api--order-service/internal/api/protos/invoice"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InvoiceClient struct {
	client invoice.InvoicerClient
	conn   *grpc.ClientConn // Opcional: guardar para fechar depois
}

func NewInvoiceClient(addr string) *InvoiceClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Erro ao criar client invoice", err)
		return nil
	}
	client := invoice.NewInvoicerClient(conn)
	return &InvoiceClient{client: client, conn: conn}
}

func (pc *InvoiceClient) GetInvoice(orderID int32, ctx context.Context) (*invoice.InvoiceResponse, error) {
	req := &invoice.InvoiceRequest{OrderId: orderID}
	return pc.client.GetInvoice(ctx, req)
}

func (pc *InvoiceClient) Close() error {
	return pc.conn.Close()
}
