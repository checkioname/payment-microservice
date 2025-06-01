package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PaymentRequest struct {
	OrderId int32 `json:"order_id"`
}

type PaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type PaymentClient interface {
	PayOrder(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
}
type paymentClient struct {
	BaseURL string
	Client  *http.Client
}

func NewPaymentClient(baseURL string) PaymentClient {
	return &paymentClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

func (pc *paymentClient) PayOrder(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	url := fmt.Sprintf("%s/payment", pc.BaseURL)

	// Monta o JSON do request
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Cria a requisição HTTP
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Executa a chamada
	resp, err := pc.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica se deu erro HTTP
	if resp.StatusCode >= 300 {
		return nil, errors.New("payment service returned http error")
	}

	// Decodifica body
	var out PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
