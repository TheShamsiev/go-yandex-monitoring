package main

import (
	"fmt"
	"time"

	"go-yandex-monitoring/internal/metrics"
)

func main() {
	cfg := ParseConfig()
	fmt.Printf("[DEBUG] Server address: %s\n", cfg.ServerAddress)
	fmt.Printf("[DEBUG] Report interval: %s\n", cfg.ReportInterval)
	fmt.Printf("[DEBUG] Poll interval: %s\n", cfg.PollInterval)

	metrics := metrics.NewMetrics()

	for {
		for i := cfg.PollInterval; i < cfg.ReportInterval; i += cfg.PollInterval {
			metrics.Poll()
			time.Sleep(cfg.PollInterval)
		}
		metrics.Report(cfg.ServerAddress)
	}
}
