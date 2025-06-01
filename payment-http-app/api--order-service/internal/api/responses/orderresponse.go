package responses

type CreateOrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
