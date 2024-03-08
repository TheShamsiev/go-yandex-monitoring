package metrics

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
)

type MetricsGauge map[string]float64

type MetricsCounter map[string]int64

type Metrics struct {
	gauge   MetricsGauge
	counter MetricsCounter
}

func NewMetrics() Metrics {
	return Metrics{make(MetricsGauge), make(MetricsCounter)}
}

func (m *Metrics) Poll() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.gauge["Alloc"] = float64(memStats.Alloc)
	m.gauge["BuckHashSys"] = float64(memStats.BuckHashSys)
	m.gauge["Frees"] = float64(memStats.Frees)
	m.gauge["GCCPUFraction"] = float64(memStats.GCCPUFraction)
	m.gauge["GCSys"] = float64(memStats.GCSys)
	m.gauge["HeapAlloc"] = float64(memStats.HeapAlloc)
	m.gauge["HeapIdle"] = float64(memStats.HeapIdle)
	m.gauge["HeapInuse"] = float64(memStats.HeapInuse)
	m.gauge["HeapObjects"] = float64(memStats.HeapObjects)
	m.gauge["HeapReleased"] = float64(memStats.HeapReleased)
	m.gauge["HeapSys"] = float64(memStats.HeapSys)
	m.gauge["LastGC"] = float64(memStats.LastGC)
	m.gauge["Lookups"] = float64(memStats.Lookups)
	m.gauge["MCacheInuse"] = float64(memStats.MCacheInuse)
	m.gauge["MCacheSys"] = float64(memStats.MCacheSys)
	m.gauge["MSpanInuse"] = float64(memStats.MSpanInuse)
	m.gauge["MSpanSys"] = float64(memStats.MSpanSys)
	m.gauge["Mallocs"] = float64(memStats.Mallocs)
	m.gauge["NextGC"] = float64(memStats.NextGC)
	m.gauge["NumForcedGC"] = float64(memStats.NumForcedGC)
	m.gauge["NumGC"] = float64(memStats.NumGC)
	m.gauge["OtherSys"] = float64(memStats.OtherSys)
	m.gauge["PauseTotalNs"] = float64(memStats.PauseTotalNs)
	m.gauge["StackInuse"] = float64(memStats.StackInuse)
	m.gauge["StackSys"] = float64(memStats.StackSys)
	m.gauge["Sys"] = float64(memStats.Sys)
	m.gauge["TotalAlloc"] = float64(memStats.TotalAlloc)

	m.gauge["RandomValue"] = rand.Float64()

	m.counter["PollCount"]++
}

func sendMetric(serverAddress, metricType, metricName, value string) {
	url := fmt.Sprintf("%s/update/%s/%s/%s", serverAddress, metricType, metricName, value)

	res, err := http.DefaultClient.Post(url, "text/plain", nil)

	if err != nil {
		fmt.Printf("[ERROR] %s", err.Error())
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("[ERROR] Status code: %d", res.StatusCode)
		return
	}
}

func (m *Metrics) Report(serverAddress string) {
	for name, value := range m.gauge {
		sendMetric(serverAddress, "gauge", name, fmt.Sprint(value))
	}

	for name, value := range m.counter {
		sendMetric(serverAddress, "counter", name, fmt.Sprint(value))
	}
}
