package main

import (
	"fmt"
	"net/http"

	"go-yandex-monitoring/internal/handler"
	"go-yandex-monitoring/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	parseFlags()

	ms := storage.NewMemStorage(make(storage.MemStorageGauge), make(storage.MemStorageCounter))

	r := chi.NewRouter()

	handler.GetMetricsRouter(r, &ms)
	handler.GetMetricRouter(r, &ms)
	handler.UpdateMetricsRouter(r, &ms)

	fmt.Printf("[DEBUG] Server address: %s\n", flagAddress)
	err := http.ListenAndServe(flagAddress, r)
	if err != nil {
		panic(err)
	}
}
