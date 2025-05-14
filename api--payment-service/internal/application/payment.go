package application

import (
	"anturiocode/api--payment-service/internal/api/protos/api"
	"anturiocode/api--payment-service/internal/infrastructure/repositories"
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	max = 1.5
	min = 0.5
)

type PaymentService struct {
	R *repositories.PaymentRepository
	api.UnimplementedPayerServer
}

func NewPaymentService(r *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{R: r}
}

func (p PaymentService) ProcessPayment(ctx context.Context, request *api.PaymentRequest) (*api.PaymentResponse, error) {
	//	# Simula processamento
	n := min + rand.Float64()*(max-min)
	processingTime := time.Duration(n)

	select {
	case <-time.After(processingTime):
		success := rand.Float64() > 0.1
		if success {
			return &api.PaymentResponse{}, nil
		}
		return &api.PaymentResponse{}, fmt.Errorf("payment processing failed")

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
