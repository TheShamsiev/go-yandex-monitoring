package main

import (
	"flag"
	"os"
	"strconv"
	"time"
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

	if envServerAddress := os.Getenv("ADDRESS"); envServerAddress != "" {
		flagServerAddress = envServerAddress
	}

	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		flagServerAddress = envReportInterval
	}

	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		flagServerAddress = envPollInterval
	}
}
