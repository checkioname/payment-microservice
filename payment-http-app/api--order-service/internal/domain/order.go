package domain

type Order struct {
	Id           int32   `json:"id,omitempty"`
	CustomerId   int32   `json:"customer_id,omitempty"`
	CreationDate string  `json:"creation_date,omitempty"`
	Status       string  `json:"status,omitempty"` // e.g. 'pending', 'payment_approved', 'picking', 'invoice_issued', 'shipped', 'canceled'
	TotalAmount  float64 `json:"total_amount,omitempty"`
}
