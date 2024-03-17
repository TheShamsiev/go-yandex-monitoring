package main

import (
	"fmt"
	"net/http"
	"os"

	"go-yandex-monitoring/internal/handler"
	"go-yandex-monitoring/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := ParseConfig()

	ms := storage.NewMemStorage(make(storage.MemStorageGauge), make(storage.MemStorageCounter))

	r := chi.NewRouter()

	handler.GetMetricsRouter(r, &ms)
	handler.GetMetricRouter(r, &ms)
	handler.UpdateMetricsRouter(r, &ms)

	fmt.Printf("[DEBUG] Server address: %s\n", cfg.Address)
	err := http.ListenAndServe(cfg.Address, r)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
}
