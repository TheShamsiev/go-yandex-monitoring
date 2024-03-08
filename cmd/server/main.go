package main

import (
	"go-yandex-monitoring/internal/handler"
	"go-yandex-monitoring/internal/storage"
	"net/http"
)

func main() {
	ms := storage.NewMemStorage(make(storage.MemStorageGauge), make(storage.MemStorageCounter))

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handler.UpdateMetrics(&ms))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
