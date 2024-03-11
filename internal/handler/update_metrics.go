package handler

import (
	"go-yandex-monitoring/internal/storage"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func updateMetrics(ms storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
			ms.UpdateCounter(metricName, val)

		case "gauge":
			val, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ms.UpdateGauge(metricName, val)

		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func UpdateMetricsRouter(r chi.Router, ms storage.Storage) {
	r.Post("/update/{metric_type}/{metric_name}/{metric_value}", updateMetrics(ms))
}
