package repositories

import (
	"anturiocode/api--order-service/internal/infrastructure"
	"anturiocode/api--order-service/internal/infrastructure/repositories/queries"
	"context"

	"go.opentelemetry.io/otel/trace"
)

type OrderRepository interface {
	RegisterOrder(order_id int32, ctx context.Context) error
}

type orderRepository struct {
	s infrastructure.PostgresStore
	t trace.Tracer
}

func NewOrderRepository(store infrastructure.PostgresStore, tracer trace.Tracer) OrderRepository {
	return &orderRepository{
		s: store,
		t: tracer,
	}
}

func (o orderRepository) RegisterOrder(order_id int32, ctx context.Context) error {
	//client_id := 1
	//status := "pending"
	//total_amount := 10

	//return o.s.ExecQuery(queries.SelectOrderQuery, order_id, client_id, status, total_amount)
	return o.s.ExecQuery(queries.SelectOrderQuery)
}
