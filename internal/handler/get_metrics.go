package handler

import (
	"go-yandex-monitoring/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getMetrics(ms storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(ms.String()))
	}
}

func GetMetricsRouter(r chi.Router, ms storage.Storage) {
	r.Get("/", getMetrics(ms))
}
