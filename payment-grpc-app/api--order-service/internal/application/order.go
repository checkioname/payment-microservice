package application

import (
	"anturiocode/api--order-service/infrastructure/repositories"
	order2 "anturiocode/api--order-service/internal/api/protos/order"
	"anturiocode/api--order-service/internal/application/client"
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"math/rand"
)

type Order struct {
	OrderID    int32
	CustomerId int32
	ItemsId    []int32
}

func newOrder(customerid int32, itemlist []*order2.OrderItem) *Order {
	var itemsids []int32
	for _, obj := range itemlist {
		itemsids = append(itemsids, obj.ProductId)
	}
	return &Order{
		OrderID:    rand.Int31(),
		CustomerId: customerid,
		ItemsId:    itemsids,
	}

}

type OrderService struct {
	R         repositories.OrderRepository
	invoice   *client.InvoiceClient
	payment   *client.PaymentClient
	shipping  *client.ShippingClient
	inventory *client.InventoryClient
	t         trace.Tracer
	order2.UnimplementedOrderServiceServer
}

func NewOrderService(r repositories.OrderRepository, tracer trace.Tracer) order2.OrderServiceServer {
	return &OrderService{R: r, t: tracer, payment: &client.PaymentClient{}}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *order2.CreateOrderRequest) (*order2.CreateOrderResponse, error) {
	resp, err := o.payment.PayOrder(req)
	if err != nil {
		slog.Warn("Erro no processamento gRPC:", err)
		return &order2.CreateOrderResponse{}, errors.New("fail to create order")
	}

	if !resp.Success {
		slog.Warn("Erro no processamento do pagamento")
		return &order2.CreateOrderResponse{}, errors.New("no processamento do pagamento")
	}

	// Registra um pedido no banco e parte para o estoque
	orderDomain := newOrder(req.CustomerId, req.Items)
	fmt.Println(orderDomain)
	_ = o.R.RegisterOrder(orderDomain.OrderID, ctx) //orderId e context

	go func(orderId int32) {
		resp, err := o.invoice.GetInvoice(1234, ctx)
		if err != nil {
			slog.Error("Erro ao emitir nota fiscal", "orderId", orderId, "err", err)
			// Aqui você pode salvar numa fila de retry, marcar no banco etc
			return
		}
		slog.Info("Nota fiscal emitida com sucesso", "orderId", orderId, "invoiceNumber", resp.OrderId)
	}(orderDomain.OrderID)

	// Separa o produto no estoque (chama API de estoque) e chama API de nota fiscal numa go routine
	o.inventory.CheckAndReserveStock(orderDomain.OrderID)

	// Chama API de logistica para iniciar entrega (devolver codigo de rastreio
	o.shipping.ShipOrder(orderDomain.OrderID)

	return &order2.CreateOrderResponse{}, nil
}

func (o *OrderService) UpdateOrderStatus(context.Context, *order2.UpdateOrderStatusRequest) (*order2.UpdateOrderStatusResponse, error) {
	return &order2.UpdateOrderStatusResponse{}, nil
}
