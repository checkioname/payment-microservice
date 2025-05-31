package application

import (
	shipping "anturiocode/api--logistics-service/internal/api/protos"
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"go.opentelemetry.io/otel/trace"

	"time"
)

type ShippingService struct {
	tracer trace.Tracer
	shipping.UnimplementedShippingServiceServer
}

func NewShippingService(t trace.Tracer) shipping.ShippingServiceServer {
	return &ShippingService{tracer: t}
}

type FakeShippingResponse struct {
	Success      bool
	TrackingCode string `faker:"uuid_digit"`
	Message      string
}

func (i ShippingService) DispatchOrder(context.Context, *shipping.DispatchOrderRequest) (*shipping.DispatchOrderResponse, error) {
	fmt.Println("Dispatch called")
	time.Sleep(1 * time.Second)

	a := FakeShippingResponse{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	return &shipping.DispatchOrderResponse{Success: true, TrackingCode: a.TrackingCode, Message: "All good"}, nil
}
