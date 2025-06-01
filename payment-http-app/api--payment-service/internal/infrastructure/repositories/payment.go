package repositories

import (
	"anturiocode/api--order-service/internal/infrastructure"
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

type PaymentRepository interface {
	RegisterPayment(price float32, product_id int32, ctx context.Context) error
}

type paymentRepository struct {
	s infrastructure.PostgresStore
	t trace.Tracer
}

func NewPaymentRepository(store infrastructure.PostgresStore, tracer trace.Tracer) PaymentRepository {
	return &paymentRepository{
		s: store,
		t: tracer,
	}
}

func (p paymentRepository) RegisterPayment(price float32, product_id int32, ctx context.Context) error {
	_, span := p.t.Start(ctx, "RegisterPayment Database")
	defer span.End()

	time.Sleep(2 * time.Second)
	//p.s.Query(queries.RegisterPaymentQuery)
	return nil
}
