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

func TestUpdateMetricsRouter(t *testing.T) {
	ms := storage.NewMemStorage(storage.MemStorageGauge{"a": 1.0}, storage.MemStorageCounter{"b": 2})
	r := chi.NewRouter()
	h := Handler{&ms}
	h.UpdateMetrics(r)
	srv := httptest.NewServer(r)
	defer srv.Close()

	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "invalid method type (GET)",
			method:         "GET",
			url:            "/update/gauge/a/2.0",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid method type (PUT)",
			method:         "PUT",
			url:            "/update/counter/b/3",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid method type (DELETE)",
			method:         "DELETE",
			url:            "/update/gauge/a/2.0",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "valid reqeust for metric type 'counter'",
			method:         "POST",
			url:            "/update/counter/b/3",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "valid reqeust for metric type 'gauge'",
			method:         "POST",
			url:            "/update/gauge/a/2.0",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "request without metric name",
			method:         "POST",
			url:            "/update/counter/3",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "request with invalid metric type",
			method:         "POST",
			url:            "/update/amogus/gauge/2.0",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "request with invalid metric value (counter)",
			method:         "POST",
			url:            "/update/counter/b/3.14",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "request with invalid metric value (gauge)",
			method:         "POST",
			url:            "/update/gauge/a/pi",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = test.method
			req.URL = srv.URL + test.url

			res, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, test.expectedStatus, res.StatusCode(), "Wrong response status code")
		})
	}
}
