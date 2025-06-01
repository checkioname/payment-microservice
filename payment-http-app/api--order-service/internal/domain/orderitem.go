package domain

type OrderItem struct {
	ProductId int32   `json:"product_id,omitempty"`
	Quantity  int32   `json:"quantity,omitempty"`
	UnitPrice float64 `json:"unit_price,omitempty"`
}
