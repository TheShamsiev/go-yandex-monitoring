package storage

import (
	"testing"
)

func TestGauge(t *testing.T) {
	tests := []struct {
		name           string
		gauge          map[string]float64
		key            string
		expected_value float64
		expected_error error
	}{
		{
			name:           "empty gauge",
			gauge:          make(map[string]float64),
			key:            "some_key",
			expected_value: 0,
			expected_error: KeyNotFoundError{"some_key"},
		},
		{
			name:           "no key in gauge",
			gauge:          map[string]float64{"pi": 3.14, "answer": 42},
			key:            "some_key",
			expected_value: 0,
			expected_error: KeyNotFoundError{"some_key"},
		},
		{
			name:           "key in gauge",
			gauge:          map[string]float64{"pi": 3.14, "answer": 42},
			key:            "answer",
			expected_value: 42,
			expected_error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter := make(map[string]int64)
			ms := MemStorage{test.gauge, counter}

			actual_value, actual_error := ms.Gauge(test.key)
			if actual_value != test.expected_value || actual_error != test.expected_error {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actual_value,
					actual_error,
					test.expected_value,
					test.expected_error,
				)
			}
		})
	}
}

func TestCounter(t *testing.T) {
	tests := []struct {
		name           string
		counter        map[string]int64
		key            string
		expected_value int64
		expected_error error
	}{
		{
			name:           "empty counter",
			counter:        make(map[string]int64),
			key:            "some_key",
			expected_value: 0,
			expected_error: KeyNotFoundError{"some_key"},
		},
		{
			name:           "no key in counter",
			counter:        map[string]int64{"answer": 42, "false": 0},
			key:            "question",
			expected_value: 0,
			expected_error: KeyNotFoundError{"question"},
		},
		{
			name:           "key in counter",
			counter:        map[string]int64{"answer": 42, "true": 1},
			key:            "true",
			expected_value: 1,
			expected_error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gauge := make(map[string]float64)
			ms := MemStorage{gauge, test.counter}

			actual_value, actual_error := ms.Counter(test.key)
			if actual_value != test.expected_value || actual_error != test.expected_error {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actual_value,
					actual_error,
					test.expected_value,
					test.expected_error,
				)
			}
		})
	}
}

func TestUpdateGauge(t *testing.T) {
	type key_value_pair struct {
		key string
		val float64
	}

	tests := []struct {
		name           string
		gauge          map[string]float64
		input_sequence []key_value_pair
		key            string
		expected_value float64
		expected_error error
	}{
		{
			name:           "entry has not been passed to storage",
			gauge:          map[string]float64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:            "e",
			expected_value: 0,
			expected_error: KeyNotFoundError{"e"},
		},
		{
			name:           "entry has been passed to storage one time",
			gauge:          map[string]float64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:            "c",
			expected_value: 3,
			expected_error: nil,
		},
		{
			name:           "entry has been passed to storage multiple times",
			gauge:          map[string]float64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}, {"c", 10}},
			key:            "c",
			expected_value: 10,
			expected_error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter := make(map[string]int64)
			ms := MemStorage{test.gauge, counter}

			for _, kv := range test.input_sequence {
				ms.UpdateGauge(kv.key, kv.val)
			}

			actual_value, actual_error := ms.Gauge(test.key)
			if actual_value != test.expected_value || actual_error != test.expected_error {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actual_value,
					actual_error,
					test.expected_value,
					test.expected_error,
				)
			}
		})
	}
}

func TestUpdateCounter(t *testing.T) {
	type key_value_pair struct {
		key string
		val int64
	}

	tests := []struct {
		name           string
		counter        map[string]int64
		input_sequence []key_value_pair
		key            string
		expected_value int64
		expected_error error
	}{
		{
			name:           "entry has not been passed to storage",
			counter:        map[string]int64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:            "e",
			expected_value: 0,
			expected_error: KeyNotFoundError{"e"},
		},
		{
			name:           "entry has been passed to storage one time",
			counter:        map[string]int64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}},
			key:            "c",
			expected_value: 3,
			expected_error: nil,
		},
		{
			name:           "entry has been passed to storage multiple times",
			counter:        map[string]int64{"a": 1},
			input_sequence: []key_value_pair{{"b", 2}, {"c", 3}, {"d", 4}, {"c", 10}},
			key:            "c",
			expected_value: 13,
			expected_error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gauge := make(map[string]float64)
			ms := MemStorage{gauge, test.counter}

			for _, kv := range test.input_sequence {
				ms.UpdateCounter(kv.key, kv.val)
			}

			actual_value, actual_error := ms.Counter(test.key)
			if actual_value != test.expected_value || actual_error != test.expected_error {
				t.Errorf(
					"actual_value: %v, actual_error: %v; expected_value: %v, expected_error: %v",
					actual_value,
					actual_error,
					test.expected_value,
					test.expected_error,
				)
			}
		})
	}
}
