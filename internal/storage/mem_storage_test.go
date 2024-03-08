package storage

import (
	"testing"
)

func TestGauge(t *testing.T) {
	tests := []struct {
		name          string
		gauge         MemStorageGauge
		key           string
		expectedValue float64
		expectedError error
	}{
		{
			name:          "empty gauge",
			gauge:         make(MemStorageGauge),
			key:           "some_key",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"some_key"},
		},
		{
			name:          "no key in gauge",
			gauge:         MemStorageGauge{"pi": 3.14, "answer": 42},
			key:           "some_key",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"some_key"},
		},
		{
			name:          "key in gauge",
			gauge:         MemStorageGauge{"pi": 3.14, "answer": 42},
			key:           "answer",
			expectedValue: 42,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter := make(MemStorageCounter)
			ms := NewMemStorage(test.gauge, counter)

			actualValue, actualError := ms.Gauge(test.key)
			if actualValue != test.expectedValue || actualError != test.expectedError {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actualValue,
					actualError,
					test.expectedValue,
					test.expectedError,
				)
			}
		})
	}
}

func TestCounter(t *testing.T) {
	tests := []struct {
		name          string
		counter       MemStorageCounter
		key           string
		expectedValue int64
		expectedError error
	}{
		{
			name:          "empty counter",
			counter:       make(MemStorageCounter),
			key:           "some_key",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"some_key"},
		},
		{
			name:          "no key in counter",
			counter:       MemStorageCounter{"answer": 42, "false": 0},
			key:           "question",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"question"},
		},
		{
			name:          "key in counter",
			counter:       MemStorageCounter{"answer": 42, "true": 1},
			key:           "true",
			expectedValue: 1,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gauge := make(MemStorageGauge)
			ms := NewMemStorage(gauge, test.counter)

			actualValue, actualError := ms.Counter(test.key)
			if actualValue != test.expectedValue || actualError != test.expectedError {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actualValue,
					actualError,
					test.expectedValue,
					test.expectedError,
				)
			}
		})
	}
}

func TestUpdateGauge(t *testing.T) {
	type keyValuePair struct {
		key string
		val float64
	}

	tests := []struct {
		name          string
		gauge         MemStorageGauge
		inputSequence []keyValuePair
		key           string
		expectedValue float64
		expectedError error
	}{
		{
			name:          "entry has not been passed to storage",
			gauge:         MemStorageGauge{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:           "e",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"e"},
		},
		{
			name:          "entry has been passed to storage one time",
			gauge:         MemStorageGauge{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:           "c",
			expectedValue: 3,
			expectedError: nil,
		},
		{
			name:          "entry has been passed to storage multiple times",
			gauge:         MemStorageGauge{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}, {"c", 10}},
			key:           "c",
			expectedValue: 10,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter := make(MemStorageCounter)
			ms := NewMemStorage(test.gauge, counter)

			for _, kv := range test.inputSequence {
				ms.UpdateGauge(kv.key, kv.val)
			}

			actualValue, actualError := ms.Gauge(test.key)
			if actualValue != test.expectedValue || actualError != test.expectedError {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actualValue,
					actualError,
					test.expectedValue,
					test.expectedError,
				)
			}
		})
	}
}

func TestUpdateCounter(t *testing.T) {
	type keyValuePair struct {
		key string
		val int64
	}

	tests := []struct {
		name          string
		counter       MemStorageCounter
		inputSequence []keyValuePair
		key           string
		expectedValue int64
		expectedError error
	}{
		{
			name:          "entry has not been passed to storage",
			counter:       MemStorageCounter{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:           "e",
			expectedValue: 0,
			expectedError: KeyNotFoundError{"e"},
		},
		{
			name:          "entry has been passed to storage one time",
			counter:       MemStorageCounter{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:           "c",
			expectedValue: 3,
			expectedError: nil,
		},
		{
			name:          "entry has been passed to storage multiple times",
			counter:       MemStorageCounter{"a": 1},
			inputSequence: []keyValuePair{{"b", 2}, {"c", 3}, {"d", 4}, {"c", 10}},
			key:           "c",
			expectedValue: 13,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gauge := make(MemStorageGauge)
			ms := NewMemStorage(gauge, test.counter)

			for _, kv := range test.inputSequence {
				ms.UpdateCounter(kv.key, kv.val)
			}

			actualValue, actualError := ms.Counter(test.key)
			if actualValue != test.expectedValue || actualError != test.expectedError {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actualValue,
					actualError,
					test.expectedValue,
					test.expectedError,
				)
			}
		})
	}
}
