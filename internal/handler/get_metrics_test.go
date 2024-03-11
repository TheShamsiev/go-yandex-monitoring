package handler

import (
	"go-yandex-monitoring/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetMetricsRouter(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		ms             storage.MemStorage
		expectedBody   string
		expectedStatus int
	}{
		{
			name:   "empty storage",
			method: "GET",
			ms: storage.NewMemStorage(
				storage.MemStorageGauge{},
				storage.MemStorageCounter{},
			),
			expectedBody:   "",
			expectedStatus: http.StatusOK,
		},
		{
			name:   "empty storage",
			method: "GET",
			ms: storage.NewMemStorage(
				storage.MemStorageGauge{"a": 3.14},
				storage.MemStorageCounter{"b": 42},
			),
			expectedBody:   "a: 3.14\nb: 42",
			expectedStatus: http.StatusOK,
		},
		{
			name:   "invalid method (POST)",
			method: "POST",
			ms: storage.NewMemStorage(
				storage.MemStorageGauge{"a": 3.14},
				storage.MemStorageCounter{"b": 42},
			),
			expectedBody:   "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "invalid method (PUT)",
			method: "PUT",
			ms: storage.NewMemStorage(
				storage.MemStorageGauge{"a": 3.14},
				storage.MemStorageCounter{"b": 42},
			),
			expectedBody:   "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "invalid method (DELETE)",
			method: "DELETE",
			ms: storage.NewMemStorage(
				storage.MemStorageGauge{"a": 3.14},
				storage.MemStorageCounter{"b": 42},
			),
			expectedBody:   "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter()
			GetMetricsRouter(r, &test.ms)
			srv := httptest.NewServer(r)
			defer srv.Close()

			req := resty.New().R()
			req.Method = test.method
			req.URL = srv.URL

			res, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, test.expectedStatus, res.StatusCode(), "Wrong response status code")
			if test.expectedStatus != http.StatusOK {
				assert.Equal(t, test.expectedBody, string(res.Body()), "Wrong response body")
			}
		})
	}
}
