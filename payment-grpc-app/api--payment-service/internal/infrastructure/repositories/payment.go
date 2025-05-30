package repositories

import (
	"anturiocode/api--payment-service/internal/infrastructure"
	"anturiocode/api--payment-service/internal/infrastructure/repositories/queries"
	"context"
	"fmt"
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

	rows, err := p.s.Query(queries.RegisterPaymentQuery)
	if err != nil {
		fmt.Println("falhei aqui")
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		
		fmt.Println(rows)
	}

	return nil
}
