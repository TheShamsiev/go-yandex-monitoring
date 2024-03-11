package main

import (
	"fmt"
	"time"

	"go-yandex-monitoring/internal/metrics"
)

func main() {
	parseFlags()
	fmt.Printf("[DEBUG] Server address: %s\n", flagServerAddress)
	fmt.Printf("[DEBUG] Report interval: %s\n", flagReportInterval)
	fmt.Printf("[DEBUG] Poll interval: %s\n", flagPollInterval)

	metrics := metrics.NewMetrics()

	for {
		for i := flagPollInterval; i < flagReportInterval; i += flagPollInterval {
			metrics.Poll()
			time.Sleep(flagPollInterval)
		}
		metrics.Report(flagServerAddress)
	}
}
