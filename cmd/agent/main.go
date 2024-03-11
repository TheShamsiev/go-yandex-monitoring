package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"go-yandex-monitoring/internal/metrics"
)

var (
	flagServerAddress  string
	flagReportInterval = 10 * time.Second
	flagPollInterval   = 2 * time.Second
)

func parseSeconds(d *time.Duration) func(string) error {
	return func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		*d = time.Duration(n) * time.Second
		return nil
	}
}

func parseFlags() {
	flag.StringVar(&flagServerAddress, "a", "localhost:8080", "address of a metrics server")
	flag.Func("r", "metrics report interval", parseSeconds(&flagReportInterval))
	flag.Func("p", "metrics poll interval", parseSeconds(&flagPollInterval))
	flag.Parse()
}

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
