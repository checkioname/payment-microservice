package responses

type CheckAndReserveStockResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
