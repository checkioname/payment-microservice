package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type InvoiceClient interface {
	GetInvoice(ctx context.Context, orderID int32) (*InvoiceResponse, error)
}

type invoiceClient struct {
	BaseURL string
	Client  *http.Client
}

func NewInvoiceClient(baseURL string) InvoiceClient {
	return &invoiceClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

// Ajuste conforme o contrato real da sua API REST
type InvoiceRequest struct {
	OrderId int32 `json:"order_id"`
}

type InvoiceResponse struct {
	Success   bool   `json:"success"`
	OrderId   int32  `json:"order_id,omitempty"`
	InvoiceId int32  `json:"invoice_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (ic *invoiceClient) GetInvoice(ctx context.Context, orderID int32) (*InvoiceResponse, error) {
	reqBody := InvoiceRequest{OrderId: orderID}
	url := fmt.Sprintf("%s/invoices", ic.BaseURL)
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
		return nil, errors.New("invoice service returned http error")
	}

	var out InvoiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
