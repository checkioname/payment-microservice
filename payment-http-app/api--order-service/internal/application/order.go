package application

import (
	"anturiocode/api--order-service/internal/api/requests"
	"anturiocode/api--order-service/internal/api/responses"
	"anturiocode/api--order-service/internal/application/client"
	"anturiocode/api--order-service/internal/domain"
	"anturiocode/api--order-service/internal/infrastructure/observability"
	"anturiocode/api--order-service/internal/infrastructure/repositories"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"

	"go.opentelemetry.io/otel/trace"
)

type Order struct {
	OrderID    int32
	CustomerId int32
	ItemsId    []int32
}

func newOrder(customerid int32, itemlist []*domain.OrderItem) *Order {
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

type OrderService interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) (*responses.CreateOrderResponse, error)
}

type orderService struct {
	R         repositories.OrderRepository
	payment   client.PaymentClient
	invoice   client.InvoiceClient
	inventory client.InventoryClient
	shipping  client.ShippingClient
	tracer    trace.Tracer
	metrics   observability.Metrics
}

func NewOrderService(r repositories.OrderRepository, tracer trace.Tracer, m *observability.Metrics) OrderService {
	return &orderService{
		R:         r,
		payment:   client.NewPaymentClient("payment-service:8008"),
		inventory: client.NewInventoryClient("inventory-service:8010"),
		shipping:  client.NewShippingClient("shipping-service:8011"),
		invoice:   client.NewInvoiceClient("invoice-service:8012"),
		tracer:    tracer,
		metrics:   *m,
	}
}

func (o *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) (*responses.CreateOrderResponse, error) {
	resp, err := o.payment.PayOrder(ctx, &client.PaymentRequest{OrderId: req.CustomerId})

	fmt.Println("pagamento request recebi")
	if err != nil {
		slog.Warn("Erro no pagamento http:", err)
		return &responses.CreateOrderResponse{Order: &domain.Order{Status: "Erro interno"}}, nil
	}

	if !resp.Success {
		slog.Warn("Erro no processamento do pagamento")
		return &responses.CreateOrderResponse{Order: &domain.Order{Status: "Erro no pagamento"}}, errors.New("no processamento do pagamento")
	}

	// Registra um pedido no banco e parte para o estoque
	orderDomain := newOrder(11111, req.Items)
	_ = o.R.RegisterOrder(orderDomain.OrderID, ctx)
	fmt.Println("registrei ordem no banco")

	go func(orderId int32) {
		resp, err := o.invoice.GetInvoice(ctx, 1234)
		if err != nil {
			slog.Error("Erro ao emitir nota fiscal", "orderId", orderId, "err", err)
			return
		}
		slog.Info("Nota fiscal emitida com sucesso", "orderId", orderId, "invoiceNumber", resp.OrderId)
	}(orderDomain.OrderID)

	stockResp, err := o.inventory.CheckAndReserveStock(ctx, orderDomain.OrderID)
	if !stockResp.Success {

		return &responses.CreateOrderResponse{Order: &domain.Order{Status: "Erro no estoque"}}, errors.New("erro na separacao do pedido")
	}

	shipResp, err := o.shipping.ShipOrder(ctx, orderDomain.OrderID)
	if !shipResp.Success {
		return &responses.CreateOrderResponse{Order: &domain.Order{Status: "Erro na logistica"}}, errors.New("erro na logistica")
	}

	return &responses.CreateOrderResponse{Order: &domain.Order{Status: "Pedido criado"}}, nil
}
