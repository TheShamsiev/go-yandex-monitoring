package storage

import (
	"fmt"
)

type Storage interface {
	Gauge(key string) (float64, error)
	Counter(key string) (int64, error)
	UpdateGauge(key string, val float64)
	UpdateCounter(key string, val int64)
	String() string
}

type KeyNotFoundError struct {
	Key string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key not found: %s", e.Key)
}
