package api

import (
	"anturiocode/api--inventory-service/internal/api/requests"
	"anturiocode/api--inventory-service/internal/application"
	"github.com/go-chi/chi/v5"

	"encoding/json"
	"log"
	"net/http"
)

type InventoryHandler struct {
	R *chi.Mux
	S application.InventoryService
}

func (a InventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.R.ServeHTTP(w, r)
}

func SendJson(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(rawData)
}

// RESERVE ORDER ITEMS
/////////////////

func (a InventoryHandler) HandleCheckAndReserveStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse do body
	var client requests.CheckAndReserveStockRequest

	// fmt.Printf("MENSAGEM RECEBIDA DA CADASTRAR REQUEST: %v \n", r.Body)
	resp, err := a.S.CheckAndReserveStock(ctx, &client)
	if err != nil {
		log.Printf("Nao foi possivel criar o cliente %v", err)
		SendJson(w, 400, "Estoque com falha")
		return
	}

	SendJson(w, 200, resp)
}
