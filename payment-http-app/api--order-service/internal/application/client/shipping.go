package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ShippingClient interface {
	ShipOrder(ctx context.Context, id int32) (*DispatchOrderResponse, error)
}

type shippingClient struct {
	BaseURL string
	Client  *http.Client
}

func NewShippingClient(baseURL string) ShippingClient {
	return &shippingClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

// Adapte estes structs conforme o contrato do seu microserviço HTTP de shipping
type DispatchOrderRequest struct {
	OrderId int32 `json:"order_id"`
}

type DispatchOrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// ShipOrder faz a chamada HTTP para despachar o pedido
func (sc *shippingClient) ShipOrder(ctx context.Context, id int32) (*DispatchOrderResponse, error) {
	reqBody := DispatchOrderRequest{OrderId: id}
	url := fmt.Sprintf("%s/shipping/dispatch", sc.BaseURL) // Ajude esta rota conforme sua API REST

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := sc.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, errors.New("shipping service returned http error")
	}

	var out DispatchOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
