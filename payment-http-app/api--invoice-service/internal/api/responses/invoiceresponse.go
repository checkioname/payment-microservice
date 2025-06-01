package responses

type InvoiceResponse struct {
	Success   bool   `json:"success"`
	OrderId   int32  `json:"order_id,omitempty"`
	InvoiceId int32  `json:"invoice_id,omitempty"`
	Message   string `json:"message,omitempty"`
}
