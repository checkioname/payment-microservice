package client

import (
	"anturiocode/api--order-service/internal/api/protos/Inventory"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	client inventory.InventoryServiceClient
	conn   *grpc.ClientConn
}

func NewInventoryClient(addr string) (*InventoryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := inventory.NewInventoryServiceClient(conn)
	return &InventoryClient{client: client, conn: conn}, nil
}

func (pc *InventoryClient) CheckAndReserveStock(id int32) (*inventory.CheckAndReserveStockResponse, error) {
	req := &inventory.CheckAndReserveStockRequest{OrderId: id}
	return pc.client.CheckAndReserveStock(context.Background(), req)
}

func (pc *InventoryClient) Close() error {
	return pc.conn.Close()
}
