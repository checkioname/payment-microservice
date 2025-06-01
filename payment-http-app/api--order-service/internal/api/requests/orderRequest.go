package requests

import "anturiocode/api--order-service/internal/domain"

type CreateOrderRequest struct {
	CustomerId int32               `json:"customer_id,omitempty"`
	Items      []*domain.OrderItem `json:"items,omitempty"`
}
