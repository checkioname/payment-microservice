package repositories

import (
	"anturiocode/api--order-service/infrastructure"
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
	//TODO implement me
	panic("implement me")
}
