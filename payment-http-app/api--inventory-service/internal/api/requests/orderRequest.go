package requests

type CheckAndReserveStockRequest struct {
	OrderId int32 `json:"order_id"`
}
