package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) updateMetrics(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metric_type")
	metricName := chi.URLParam(r, "metric_name")
	metricValue := chi.URLParam(r, "metric_value")

	switch metricType {
	case "counter":
		val, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.Storage.UpdateCounter(metricName, val)

	case "gauge":
		val, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.Storage.UpdateGauge(metricName, val)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) UpdateMetrics(r chi.Router) {
	r.Post("/update/{metric_type}/{metric_name}/{metric_value}", h.updateMetrics)
}
