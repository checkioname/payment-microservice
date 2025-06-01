package application

import (
	"anturiocode/api--logistics-service/internal/api/requests"
	"anturiocode/api--logistics-service/internal/api/responses"
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"go.opentelemetry.io/otel/trace"

	"time"
)

type ShippingService interface {
	DispatchOrder(context.Context, *requests.DispatchOrderRequest) (*responses.DispatchOrderResponse, error)
}

type shippingService struct {
	tracer trace.Tracer
}

func NewShippingService(t trace.Tracer) ShippingService {
	return &shippingService{tracer: t}
}

type FakeShippingResponse struct {
	Success      bool
	TrackingCode string `faker:"uuid_digit"`
	Message      string
}

func (i *shippingService) DispatchOrder(context.Context, *requests.DispatchOrderRequest) (*responses.DispatchOrderResponse, error) {
	fmt.Println("Dispatch called")
	time.Sleep(1 * time.Second)

	a := FakeShippingResponse{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	return &responses.DispatchOrderResponse{Success: true, TrackingCode: a.TrackingCode, Message: "All good"}, nil
}
