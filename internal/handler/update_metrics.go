package handler

import (
	"go-yandex-monitoring/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func UpdateMetrics(ms storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := strings.Split(r.URL.Path, "/")
		if len(path) != 5 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metric_type := path[len(path)-3]
		metric_name := path[len(path)-2]
		metric_value := path[len(path)-1]

		switch metric_type {
		case "counter":
			counter, err := strconv.ParseInt(metric_value, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ms.UpdateCounter(metric_name, counter)
		case "gauge":
			gauge, err := strconv.ParseFloat(metric_value, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ms.UpdateGauge(metric_name, gauge)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
