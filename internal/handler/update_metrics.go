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

		metricType := path[len(path)-3]
		metricName := path[len(path)-2]
		metricValue := path[len(path)-1]

		switch metricType {
		case "counter":
			counter, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ms.UpdateCounter(metricName, counter)
		case "gauge":
			gauge, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ms.UpdateGauge(metricName, gauge)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
