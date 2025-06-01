package api

import (
	"anturiocode/api--logistics-service/internal/api/requests"
	"anturiocode/api--logistics-service/internal/application"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type ShippingHandler struct {
	R *chi.Mux
	S application.ShippingService
}

func (a *ShippingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.R.ServeHTTP(w, r)
}

func SendJson(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(rawData)
}

// DISPATCH ORDER ITEMS
/////////////////

func (a *ShippingHandler) HandleDispatchOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse do body
	var client requests.DispatchOrderRequest

	// fmt.Printf("MENSAGEM RECEBIDA DA CADASTRAR REQUEST: %v \n", r.Body)
	resp, err := a.S.DispatchOrder(ctx, &client)
	if err != nil {
		log.Printf("Nao foi possivel enviar o pedido %v", err)
		SendJson(w, 400, "logistica com falha")
		return
	}

	SendJson(w, 200, resp)
}
