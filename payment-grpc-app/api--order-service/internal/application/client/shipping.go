package client

import (
	"anturiocode/api--order-service/internal/api/protos/shipping"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShippingClient struct {
	client shipping.ShippingServiceClient
	conn   *grpc.ClientConn
}

func NewShippingClient(addr string) (*ShippingClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := shipping.NewShippingServiceClient(conn)
	return &ShippingClient{client: client, conn: conn}, nil
}

func (pc *ShippingClient) ShipOrder(id int32) (*shipping.DispatchOrderResponse, error) {
	req := &shipping.DispatchOrderRequest{OrderId: id}
	return pc.client.DispatchOrder(context.Background(), req)
}

func (pc *ShippingClient) Close() error {
	return pc.conn.Close()
}
