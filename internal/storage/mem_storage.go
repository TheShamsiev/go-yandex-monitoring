package storage

import (
	"fmt"
	"strings"
)

type MemStorageGauge map[string]float64

type MemStorageCounter map[string]int64

type MemStorage struct {
	gauge   MemStorageGauge
	counter MemStorageCounter
}

func NewMemStorage(gauge MemStorageGauge, counter MemStorageCounter) MemStorage {
	return MemStorage{gauge, counter}
}

func (ms *MemStorage) Gauge(key string) (float64, error) {
	val, ok := ms.gauge[key]
	if !ok {
		return 0, KeyNotFoundError{key}
	}
	return val, nil
}

func (ms *MemStorage) Counter(key string) (int64, error) {
	val, ok := ms.counter[key]
	if !ok {
		return 0, KeyNotFoundError{key}
	}
	return val, nil
}

func (ms *MemStorage) UpdateGauge(key string, val float64) {
	ms.gauge[key] = val
}

func (ms *MemStorage) UpdateCounter(key string, val int64) {
	ms.counter[key] += val
}

func (ms *MemStorage) String() string {
	var metrics []string

	for k, v := range ms.gauge {
		metrics = append(metrics, fmt.Sprintf("%s: %v", k, v))
	}

	for k, v := range ms.counter {
		metrics = append(metrics, fmt.Sprintf("%s: %v", k, v))
	}

	return strings.Join(metrics, "\n")
}
