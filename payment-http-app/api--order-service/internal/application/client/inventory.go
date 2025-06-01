package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type InventoryClient interface {
	CheckAndReserveStock(ctx context.Context, id int32) (*CheckAndReserveStockResponse, error)
}

type inventoryClient struct {
	BaseURL string
	Client  *http.Client
}

func NewInventoryClient(baseURL string) InventoryClient {
	return &inventoryClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

type CheckAndReserveStockRequest struct {
	OrderId int32 `json:"order_id"`
}

type CheckAndReserveStockResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func (ic *inventoryClient) CheckAndReserveStock(ctx context.Context, id int32) (*CheckAndReserveStockResponse, error) {
	reqBody := CheckAndReserveStockRequest{OrderId: id}

	url := fmt.Sprintf("%s/inventory", ic.BaseURL)
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := ic.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, errors.New("inventory service returned http error")
	}

	var out CheckAndReserveStockResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
