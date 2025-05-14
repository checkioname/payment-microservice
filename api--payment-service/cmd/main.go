package main

import (
	"anturiocode/api--payment-service/internal/api/protos/api"
	"anturiocode/api--payment-service/internal/application"
	"anturiocode/api--payment-service/internal/infrastructure"
	"anturiocode/api--payment-service/internal/infrastructure/config"
	"anturiocode/api--payment-service/internal/infrastructure/repositories"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 8008, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	// dependencies
	cfg, err := config.LoadConfig("config.dev.json")
	grpcServer := grpc.NewServer(opts...)

	// IOC
	store, _ := infrastructure.NewPostgresStore(cfg.Database)
	repo := repositories.NewPaymentRepository(store)

	api.RegisterPayerServer(grpcServer, application.NewPaymentService(&repo))
	grpcServer.Serve(lis)

}
