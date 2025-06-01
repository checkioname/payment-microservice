package api

import (
	"anturiocode/api--order-service/internal/api/requests"
	"anturiocode/api--order-service/internal/application"
	"github.com/go-chi/chi/v5"

	"encoding/json"
	"log"
	"net/http"
)

type OrderHandler struct {
	R *chi.Mux
	S application.OrderService
}

func (a OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.R.ServeHTTP(w, r)
}

func SendJson(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(rawData)
}

// CREATE ORDER
/////////////////

func (a OrderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// parse do body
	var client requests.CreateOrderRequest
	// fmt.Printf("MENSAGEM RECEBIDA DA CADASTRAR REQUEST: %v \n", r.Body)
	result, err := a.S.CreateOrder(ctx, &client)
	if err != nil {
		log.Printf("Nao foi possivel criar o cliente %v", err)
	}

	SendJson(w, 200, result)
}
