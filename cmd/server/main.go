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
	h := handler.Handler{Storage: &ms}
	r := chi.NewRouter()

	h.GetMetric(r)
	h.GetMetrics(r)
	h.UpdateMetrics(r)

	fmt.Printf("[DEBUG] Server address: %s\n", cfg.Address)
	err := http.ListenAndServe(cfg.Address, r)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
}
