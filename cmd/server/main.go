package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go-yandex-monitoring/internal/handler"
	"go-yandex-monitoring/internal/storage"
)

func main() {
	ms := storage.NewMemStorage(make(storage.MemStorageGauge), make(storage.MemStorageCounter))

	r := chi.NewRouter()

	handler.GetMetricsRouter(r, &ms)
	handler.GetMetricRouter(r, &ms)
	handler.UpdateMetricsRouter(r, &ms)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
