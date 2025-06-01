package main

import (
	"anturiocode/api--payment-service/internal/api"
	"anturiocode/api--payment-service/internal/application"
	"anturiocode/api--payment-service/internal/infrastructure/observability"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	//panic handler
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("PANIC capturado: %v", r)
		}
	}()

	ctx := context.Background()
	exporter, err := observability.NewOTLPExporter(ctx)
	tp := observability.NewTraceProvider(exporter)
	tracer := tp.Tracer("payment-service")

	// IOC

	if err != nil {
		fmt.Printf("Erro ao criar store Postgres: %v", err)
	}

	app := application.NewPaymentService(tracer)
	h := api.PaymentHandler{S: app}
	handler := api.NewHandler(&h)

	// initialize server
	port := flag.Int("port", 8008, "The server port")
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal("Couldnt keep up the server alive ", err)
				panic(err)
			}
		}

	}()

	fmt.Println("Listening on port papiri", *port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
