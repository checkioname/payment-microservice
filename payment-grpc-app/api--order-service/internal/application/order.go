package application

import (
	"anturiocode/api--order-service/internal/api/protos/order"
	"anturiocode/api--order-service/internal/application/client"
	"anturiocode/api--order-service/internal/infrastructure/observability"
	"anturiocode/api--order-service/internal/infrastructure/repositories"
	"context"
	"errors"
	"log/slog"
	"math/rand"

	"go.opentelemetry.io/otel/trace"
)

type Order struct {
	OrderID    int32
	CustomerId int32
	ItemsId    []int32
}

func newOrder(customerid int32, itemlist []*order.OrderItem) *Order {
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
	metrics   *observability.Metrics
	order.UnimplementedOrderServiceServer
}

func NewOrderService(r repositories.OrderRepository, tracer trace.Tracer, m *observability.Metrics) order.OrderServiceServer {
	return &OrderService{
		R:         r,
		t:         tracer,
		metrics:   m,
		payment:   client.NewPaymentClient("8008"),
		inventory: client.NewInventoryClient(":8010"),
		shipping:  client.NewShippingClient(":8011"),
		invoice:   client.NewInvoiceClient(":8012"),
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	resp, err := o.payment.PayOrder(req)
	if err != nil {
		slog.Warn("Erro no processamento gRPC:", err)
		return &order.CreateOrderResponse{}, errors.New("fail to create order")
	}

	if !resp.Success {
		slog.Warn("Erro no processamento do pagamento")
		return &order.CreateOrderResponse{Order: &order.Order{Status: "Erro no pagamento"}}, errors.New("no processamento do pagamento")
	}

	// Registra um pedido no banco e parte para o estoque
	orderDomain := newOrder(req.CustomerId, req.Items)
	_ = o.R.RegisterOrder(orderDomain.OrderID, ctx)

	go func(orderId int32) {
		resp, err := o.invoice.GetInvoice(1234, ctx)
		if err != nil {
			slog.Error("Erro ao emitir nota fiscal", "orderId", orderId, "err", err)
			return
		}
		slog.Info("Nota fiscal emitida com sucesso", "orderId", orderId, "invoiceNumber", resp.OrderId)
	}(orderDomain.OrderID)

	stockResp, err := o.inventory.CheckAndReserveStock(orderDomain.OrderID)
	if !stockResp.Success {

		return &order.CreateOrderResponse{Order: &order.Order{Status: "Erro no estoque"}}, errors.New("erro na separacao do pedido")
	}

	shipResp, err := o.shipping.ShipOrder(orderDomain.OrderID)
	if !shipResp.Success {
		return &order.CreateOrderResponse{Order: &order.Order{Status: "Erro na logistica"}}, errors.New("erro na logistica")
	}

	return &order.CreateOrderResponse{Order: &order.Order{Status: "Pedido criado"}}, nil
}

func (o *OrderService) UpdateOrderStatus(context.Context, *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	return &order.UpdateOrderStatusResponse{}, nil
}
