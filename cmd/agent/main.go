package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
	serverAddress  = "http://localhost:8080"
)

type Metrics struct {
	MemStats    runtime.MemStats
	PollCount   int64
	RandomValue float64
}

func (m *Metrics) Poll() {
	runtime.ReadMemStats(&m.MemStats)
	m.PollCount++
	m.RandomValue = rand.Float64()
}

func sendMetric(serverAddress string, metricType string, metricName string, value any) {
	url := fmt.Sprintf("%s/update/%s/%s/%v", serverAddress, metricType, metricName, value)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Println("[ERROR]", err.Error())
	}
	req.Header.Set("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[ERROR]", err.Error())
	}
	log.Println("[DEBUG]", res.StatusCode, res.Body)
}

func (m *Metrics) Report(serverAddress string) {
	sendMetric(serverAddress, "gauge", "Alloc", m.MemStats.Alloc)
	sendMetric(serverAddress, "gauge", "BuckHashSys", m.MemStats.BuckHashSys)
	sendMetric(serverAddress, "gauge", "Frees", m.MemStats.Frees)
	sendMetric(serverAddress, "gauge", "GCCPUFraction", m.MemStats.GCCPUFraction)
	sendMetric(serverAddress, "gauge", "GCSys", m.MemStats.GCSys)
	sendMetric(serverAddress, "gauge", "HeapAlloc", m.MemStats.HeapAlloc)
	sendMetric(serverAddress, "gauge", "HeapIdle", m.MemStats.HeapIdle)
	sendMetric(serverAddress, "gauge", "HeapInuse", m.MemStats.HeapInuse)
	sendMetric(serverAddress, "gauge", "HeapObjects", m.MemStats.HeapObjects)
	sendMetric(serverAddress, "gauge", "HeapReleased", m.MemStats.HeapReleased)
	sendMetric(serverAddress, "gauge", "HeapSys", m.MemStats.HeapSys)
	sendMetric(serverAddress, "gauge", "LastGC", m.MemStats.LastGC)
	sendMetric(serverAddress, "gauge", "Lookups", m.MemStats.Lookups)
	sendMetric(serverAddress, "gauge", "MCacheInuse", m.MemStats.MCacheInuse)
	sendMetric(serverAddress, "gauge", "MCacheSys", m.MemStats.MCacheSys)
	sendMetric(serverAddress, "gauge", "MSpanInuse", m.MemStats.MSpanInuse)
	sendMetric(serverAddress, "gauge", "MSpanSys", m.MemStats.MSpanSys)
	sendMetric(serverAddress, "gauge", "Mallocs", m.MemStats.Mallocs)
	sendMetric(serverAddress, "gauge", "NextGC", m.MemStats.NextGC)
	sendMetric(serverAddress, "gauge", "NumForcedGC", m.MemStats.NumForcedGC)
	sendMetric(serverAddress, "gauge", "NumGC", m.MemStats.NumGC)
	sendMetric(serverAddress, "gauge", "OtherSys", m.MemStats.OtherSys)
	sendMetric(serverAddress, "gauge", "PauseTotalNs", m.MemStats.PauseTotalNs)
	sendMetric(serverAddress, "gauge", "StackInuse", m.MemStats.StackInuse)
	sendMetric(serverAddress, "gauge", "StackSys", m.MemStats.StackSys)
	sendMetric(serverAddress, "gauge", "Sys", m.MemStats.Sys)
	sendMetric(serverAddress, "gauge", "TotalAlloc", m.MemStats.TotalAlloc)
	sendMetric(serverAddress, "counter", "PollCount", m.PollCount)
	sendMetric(serverAddress, "gauge", "RandomValue", m.RandomValue)
}

func main() {
	var metrics Metrics

	for {
		for i := pollInterval; i < reportInterval; i += pollInterval {
			metrics.Poll()
			time.Sleep(pollInterval)
		}
		metrics.Report(serverAddress)
	}
}
