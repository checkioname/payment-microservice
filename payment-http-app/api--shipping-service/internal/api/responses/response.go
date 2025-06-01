package responses

type DispatchOrderResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message,omitempty"`
	TrackingCode string `json:"tracking_code,omitempty"`
}
