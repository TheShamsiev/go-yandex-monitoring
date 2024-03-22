package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) getMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(h.Storage.String()))
}

func (h *Handler) GetMetrics(r chi.Router) {
	r.Get("/", h.getMetrics)
}
