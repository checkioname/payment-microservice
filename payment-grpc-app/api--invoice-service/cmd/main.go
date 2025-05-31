package main

import (
	"anturiocode/api--invoice-service/internal/api/protos"
	"anturiocode/api--invoice-service/internal/application"
	"flag"
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	// IOC
	app := application.NewInvoiceService(nil)
	invoice.RegisterInvoicerServer(grpcServer, app)

	// initialize server
	port := flag.Int("port", 8012, "The server port")
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
