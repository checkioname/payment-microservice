package main

import (
	"anturiocode/api--payment-service/internal/api/protos/api"
	"anturiocode/api--payment-service/internal/application"
	"anturiocode/api--payment-service/internal/infrastructure"
	"anturiocode/api--payment-service/internal/infrastructure/config"
	"anturiocode/api--payment-service/internal/infrastructure/observability"
	"anturiocode/api--payment-service/internal/infrastructure/repositories"
	"context"
	"flag"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()

	// prometheus
	exporter, err := observability.NewOTLPExporter(ctx)
	if err != nil {
		log.Fatalf("Erro ao inicializar as métricas: %v", err)
	}

	tp := observability.NewTraceProvider(exporter)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// dependencies
	cfg, err := config.LoadConfig("config.dev.json")
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	// IOC
	store, _ := infrastructure.NewPostgresStore(cfg.Database)
	repo := repositories.NewPaymentRepository(store, observability.Tracer)

	api.RegisterPayerServer(grpcServer, application.NewPaymentService(&repo, observability.Tracer))

	// initialize server
	port := flag.Int("port", 8008, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	grpcServer.Serve(lis)

}
