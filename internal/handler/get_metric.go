package handler

import (
	"fmt"
	"go-yandex-monitoring/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getMetric(ms storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := chi.URLParam(r, "metric_type")
		metricName := chi.URLParam(r, "metric_name")

		switch metricType {
		case "counter":
			val, err := ms.Counter(metricName)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Write([]byte(fmt.Sprint(val)))

		case "gauge":
			val, err := ms.Gauge(metricName)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Write([]byte(fmt.Sprint(val)))

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func GetMetricRouter(r chi.Router, ms storage.Storage) {
	r.Get("/value/{metric_type}/{metric_name}", getMetric(ms))
}
