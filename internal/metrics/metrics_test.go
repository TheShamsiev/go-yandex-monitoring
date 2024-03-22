package metrics

import (
	"testing"
)

func checkGauge(t *testing.T, metrics *MetricsGauge, name string) {
	_, ok := (*metrics)[name]
	if !ok {
		t.Errorf("gauge metric `%s` is not present", name)
	}
}

func checkCounter(t *testing.T, metrics *MetricsCounter, name string) {
	_, ok := (*metrics)[name]
	if !ok {
		t.Errorf("counter metric `%s` is not present", name)
	}
}

func TestPoll(t *testing.T) {
	t.Run("`PollCount` increases every poll", func(t *testing.T) {
		metrics := NewMetrics()
		metrics.Poll()
		metrics.Poll()
		metrics.Poll()

		expected := int64(3)
		actual := metrics.counter["PollCount"]

		if actual != expected {
			t.Errorf("expected `PollCount` to be %d, got %d", expected, actual)
		}
	})

	t.Run("collects all required metrics", func(t *testing.T) {
		metrics := NewMetrics()
		metrics.Poll()

		checkGauge(t, &metrics.gauge, "Alloc")
		checkGauge(t, &metrics.gauge, "BuckHashSys")
		checkGauge(t, &metrics.gauge, "Frees")
		checkGauge(t, &metrics.gauge, "GCCPUFraction")
		checkGauge(t, &metrics.gauge, "GCSys")
		checkGauge(t, &metrics.gauge, "HeapAlloc")
		checkGauge(t, &metrics.gauge, "HeapIdle")
		checkGauge(t, &metrics.gauge, "HeapInuse")
		checkGauge(t, &metrics.gauge, "HeapObjects")
		checkGauge(t, &metrics.gauge, "HeapReleased")
		checkGauge(t, &metrics.gauge, "HeapSys")
		checkGauge(t, &metrics.gauge, "LastGC")
		checkGauge(t, &metrics.gauge, "Lookups")
		checkGauge(t, &metrics.gauge, "MCacheInuse")
		checkGauge(t, &metrics.gauge, "MCacheSys")
		checkGauge(t, &metrics.gauge, "MSpanInuse")
		checkGauge(t, &metrics.gauge, "MSpanSys")
		checkGauge(t, &metrics.gauge, "Mallocs")
		checkGauge(t, &metrics.gauge, "NextGC")
		checkGauge(t, &metrics.gauge, "NumForcedGC")
		checkGauge(t, &metrics.gauge, "NumGC")
		checkGauge(t, &metrics.gauge, "OtherSys")
		checkGauge(t, &metrics.gauge, "PauseTotalNs")
		checkGauge(t, &metrics.gauge, "StackInuse")
		checkGauge(t, &metrics.gauge, "StackSys")
		checkGauge(t, &metrics.gauge, "Sys")
		checkGauge(t, &metrics.gauge, "TotalAlloc")

		checkGauge(t, &metrics.gauge, "RandomValue")

		checkCounter(t, &metrics.counter, "PollCount")
	})
}
