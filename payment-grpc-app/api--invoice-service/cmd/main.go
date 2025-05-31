package main

import (
	"anturiocode/api--invoice-service/internal/infrastructure/config"
	"anturiocode/api--invoice-service/internal/infrastructure/observability"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	// prometheus
	exporter, err := observability.NewOTLPExporter(ctx)
	if err != nil {
		log.Fatalf("Erro ao inicializar as métricas: %v", err)
	}

	tp := observability.NewTraceProvider(exporter)
	tracer := tp.Tracer("myapp")
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// dependencies
	cfg, err := config.LoadConfig("config.dev.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	// IOC
	app := application.NewInventoryService(tracer)
	api.RegisterPayerServer(grpcServer, app)

	// initialize server
	port := flag.Int("port", 8008, "The server port")
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
			panic(err)
		}
		log.Printf("Servidor gRPC ouvindo em %s", lis.Addr()) // Adicione este log
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			panic(err)
		}
	}()

	fmt.Println("Listening on port", *port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
