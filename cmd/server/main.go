package main

import (
	"net/http"
	"strconv"
	"strings"
)

type Storage interface {
	UpdateGauge(key string, val float64)
	UpdateCounter(key string, val int64)
}

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func (ms *MemStorage) UpdateGauge(key string, val float64) {
	ms.gauge[key] = val
}

func (ms *MemStorage) UpdateCounter(key string, val int64) {
	ms.counter[key] += val
}

func UpdateMetrics(ms Storage) func(w http.ResponseWriter, r *http.Request) {
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

func main() {
	ms := MemStorage{make(map[string]float64), make(map[string]int64)}

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", UpdateMetrics(&ms))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
