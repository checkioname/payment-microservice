package api

import (
	"anturiocode/api--payment-service/internal/api/requests"
	"anturiocode/api--payment-service/internal/application"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type PaymentHandler struct {
	R *chi.Mux
	S application.PaymentService
}

func (a *PaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.R.ServeHTTP(w, r)
}

func SendJson(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(rawData)
}

// RESERVE ORDER ITEMS
/////////////////

func (a *PaymentHandler) HandleProcessPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse do body
	var client requests.PaymentRequest

	// fmt.Printf("MENSAGEM RECEBIDA DA CADASTRAR REQUEST: %v \n", r.Body)
	_, err := a.S.ProcessPayment(ctx, &client)
	if err != nil {
		log.Printf("Nao foi possivel nota fiscal %v", err)
		SendJson(w, 400, "nota fiscal com falha")
		return
	}

	SendJson(w, 200, "pagamento com sucesso")
}
