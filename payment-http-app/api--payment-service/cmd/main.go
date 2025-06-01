package main

import (
	"anturiocode/api--payment-service/internal/application"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

func main() {
	//reg := prometheus.NewRegistry()
	//m := observability.NewMetrics(reg)

	//panic handler
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("PANIC capturado: %v", r)
		}
	}()

	ctx := context.Background()

	// OPTL
	//exporter, err := observability.NewOTLPExporter(ctx)
	//if err != nil {
	//	fmt.Printf("Erro ao inicializar as métricas: %v", err)
	//}
	//
	//tp := observability.NewTraceProvider(exporter)
	//tracer := tp.Tracer("payment-service")
	//defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// dependencies
	cfg, err := config.LoadConfig("config.dev.json")
	if err != nil {
		fmt.Printf("Error loading config: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	// IOC
	store, err := infrastructure.NewPostgresStore(cfg.Database)
	if err != nil {
		fmt.Printf("Erro ao criar store Postgres: %v", err)
	}
	repo := repositories.NewPaymentRepository(store, tracer)
	app := application.NewPaymentService(repo, tracer, m)
	api.RegisterPayerServer(grpcServer, app)

	// initialize server
	port := flag.Int("port", 8008, "The server port")
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
			panic(err)
		}
		fmt.Printf("Servidor gRPC ouvindo em %s", lis.Addr()) // Adicione este log
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			panic(err)
		}
	}()

	fmt.Println("Listening on port papiri", *port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
