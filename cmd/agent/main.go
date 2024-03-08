package main

import (
	"time"
	"go-yandex-monitoring/internal/metrics"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
	serverAddress  = "http://localhost:8080"
)

func main() {
	metrics := metrics.NewMetrics()

	for {
		for i := pollInterval; i < reportInterval; i += pollInterval {
			metrics.Poll()
			time.Sleep(pollInterval)
		}
		metrics.Report(serverAddress)
	}
}
