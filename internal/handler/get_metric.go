package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) getMetric(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metric_type")
	metricName := chi.URLParam(r, "metric_name")

	switch metricType {
	case "counter":
		val, err := h.Storage.Counter(metricName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(fmt.Sprint(val)))

	case "gauge":
		val, err := h.Storage.Gauge(metricName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(fmt.Sprint(val)))

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) GetMetric(r chi.Router) {
	r.Get("/value/{metric_type}/{metric_name}", h.getMetric)
}
