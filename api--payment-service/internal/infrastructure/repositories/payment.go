package repositories

import "anturiocode/api--payment-service/internal/infrastructure"

type PaymentRepository interface {
	RegisterPayment(price float32, product_id int32) error
}

type paymentRepository struct {
	s infrastructure.PostgresStore
}

func NewPaymentRepository(store infrastructure.PostgresStore) PaymentRepository {
	return &paymentRepository{
		s: store,
	}
}

func (p paymentRepository) RegisterPayment(price float32, product_id int32) error {
	//TODO implement me
	panic("implement me")
}
