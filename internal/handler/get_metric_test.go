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

func TestGetMetricRouter(t *testing.T) {
	ms := storage.NewMemStorage(storage.MemStorageGauge{"a": 1.0}, storage.MemStorageCounter{"b": 2})
	r := chi.NewRouter()
	GetMetricRouter(r, &ms)
	srv := httptest.NewServer(r)
	defer srv.Close()

	tests := []struct {
		name           string
		url            string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid metric type",
			url:            "/value/joke/a",
			method:         "GET",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
		{
			name:           "invalid metric (counter) name",
			url:            "/value/counter/a",
			method:         "GET",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
		{
			name:           "invalid metric (gauge) name",
			url:            "/value/gauge/b",
			method:         "GET",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
		{
			name:           "valid metric (counter) type",
			url:            "/value/counter/b",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "2",
		},
		{
			name:           "valid metric (gauge) type",
			url:            "/value/gauge/a",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "1",
		},
		{
			name:           "invalid method type (POST)",
			url:            "/value/counter/b",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
		{
			name:           "invalid method type (PUT)",
			url:            "/value/counter/b",
			method:         "PUT",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
		{
			name:           "invalid method type (DELETE)",
			url:            "/value/counter/b",
			method:         "DELETE",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
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
			if test.expectedBody != "" {
				assert.Equal(t, test.expectedBody, string(res.Body()), "Wrong response body")
			}
		})
	}
}
