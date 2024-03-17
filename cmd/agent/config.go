package main

import (
	"flag"
	ut "go-yandex-monitoring/internal/utils/time"
	"os"
	"time"
)

type Config struct {
	ServerAddress  string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

func getTimeFlagParser(target *time.Duration, unit time.Duration) func(string) error {
	return func(s string) error {
		duration, err := ut.ParseDecimalString(s, unit)
		if err != nil {
			return err
		}
		*target = duration
		return nil
	}
}

func ParseConfig() Config {
	var (
		serverAddress  string
		reportInterval time.Duration
		pollInterval   time.Duration
	)

	flag.StringVar(&serverAddress, "a", "localhost:8080", "address of a metrics server")
	flag.Func("r", "metrics report interval", getTimeFlagParser(&reportInterval, time.Second))
	flag.Func("p", "metrics poll interval", getTimeFlagParser(&pollInterval, time.Second))
	flag.Parse()

	if envServerAddress := os.Getenv("ADDRESS"); envServerAddress != "" {
		serverAddress = envServerAddress
	}
	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		duration, err := ut.ParseDecimalString(envReportInterval, time.Second)
		if err == nil {
			reportInterval = duration
		}
	}
	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		duration, err := ut.ParseDecimalString(envPollInterval, time.Second)
		if err == nil {
			pollInterval = duration
		}
	}

	return Config{serverAddress, reportInterval, pollInterval}
}
