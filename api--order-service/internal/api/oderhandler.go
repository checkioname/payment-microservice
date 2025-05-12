package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type OrderHandler struct {
	R *chi.Mux
}

func (o OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.R.ServeHTTP(w, r)
}

func sendResponse(w http.ResponseWriter, status int, rawData any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(rawData)
}
