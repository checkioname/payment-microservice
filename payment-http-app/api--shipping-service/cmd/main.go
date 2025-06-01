package main

import (
	"anturiocode/api--logistics-service/internal/api"
	"anturiocode/api--logistics-service/internal/application"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"
	"os/signal"
)

func main() {
	// IOC
	app := application.NewShippingService(nil)
	h := api.ShippingHandler{S: app}
	handler := api.NewHandler(&h)

	// initialize server
	port := flag.Int("port", 8011, "The server port")
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
