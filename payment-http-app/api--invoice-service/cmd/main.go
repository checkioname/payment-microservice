package main

import (
	"anturiocode/api--invoice-service/internal/api"
	"anturiocode/api--invoice-service/internal/application"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"log"

	"os"
	"os/signal"
)

func main() {
	// IOC
	app := application.NewInvoiceService(nil)
	a := api.InvoiceHandler{S: app}
	handler := api.NewHandler(&a)

	// initialize server
	port := flag.Int("port", 8012, "The server port")
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal("Couldnt keep up the server alive ", err)
				panic(err)
			}
		}

	}()

	fmt.Println("Listening on port", *port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
