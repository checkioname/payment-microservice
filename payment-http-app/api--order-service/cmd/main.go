package main

import (
	"anturiocode/api--order-service/internal/api/protos/order"
	"anturiocode/api--order-service/internal/application"
	"anturiocode/api--order-service/internal/infrastructure"
	"anturiocode/api--order-service/internal/infrastructure/config"
	"anturiocode/api--order-service/internal/infrastructure/observability"
	"anturiocode/api--order-service/internal/infrastructure/repositories"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

func main() {
	reg := prometheus.NewRegistry()
	m := observability.NewMetrics(reg)

	//panic handler
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("PANIC capturado: %v", r)
		}
	}()

	ctx := context.Background()

	//prometheus
	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		err := http.ListenAndServe(":8082", nil)
		if err != nil {
			fmt.Println("Erro servidor prometheus", err.Error())
			return
		}
		fmt.Println("Servidor prometheus ouvindo em 7000")
	}()

	// OPTL``
	exporter, err := observability.NewOTLPExporter(ctx)
	if err != nil {
		fmt.Printf("Erro ao inicializar as métricas: %v", err)
	}

	tp := observability.NewTraceProvider(exporter)
	tracer := tp.Tracer("payment-service")
	defer func() { _ = tp.Shutdown(ctx) }()

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
	repo := repositories.NewOrderRepository(store, tracer)
	app := application.NewOrderService(repo, tracer, m)
	order.RegisterOrderServiceServer(grpcServer, app)

	// initialize server
	port := flag.Int("port", 8007, "The server port")
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
