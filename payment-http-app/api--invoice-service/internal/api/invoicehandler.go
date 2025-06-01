package api

import (
	"anturiocode/api--invoice-service/internal/api/requests"
	"anturiocode/api--invoice-service/internal/application"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type InvoiceHandler struct {
	R *chi.Mux
	S application.InvoiceService
}

func (a InvoiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.R.ServeHTTP(w, r)
}

func SendJson(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(rawData)
}

// RESERVE ORDER ITEMS
/////////////////

func (a InvoiceHandler) HandleGetInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse do body
	var client requests.InvoiceRequest

	// fmt.Printf("MENSAGEM RECEBIDA DA CADASTRAR REQUEST: %v \n", r.Body)
	resp, err := a.S.GetInvoice(ctx, &client)
	if err != nil {
		log.Printf("Nao foi possivel nota fiscal %v", err)
		SendJson(w, 400, "nota fiscal com falha")
		return
	}

	SendJson(w, 200, resp)
}
