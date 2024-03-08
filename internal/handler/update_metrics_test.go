package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"go-yandex-monitoring/internal/storage"
)

func TestUpdateMetrics(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		method     string
		url        string
	}{
		{
			name:       "GET request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodGet,
			url:        "/update/counter/test/1",
		},
		{
			name:       "HEAD request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodHead,
			url:        "/update/counter/test/1",
		},
		{
			name:       "PUT request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodPut,
			url:        "/update/counter/test/1",
		},
		{
			name:       "DELETE request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodDelete,
			url:        "/update/counter/test/1",
		},
		{
			name:       "CONNECT request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodConnect,
			url:        "/update/counter/test/1",
		},
		{
			name:       "OPTIONS request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodOptions,
			url:        "/update/counter/test/1",
		},
		{
			name:       "TRACE request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodTrace,
			url:        "/update/counter/test/1",
		},
		{
			name:       "PATCH request returns `http.StatusMethodNotAllowed`",
			statusCode: http.StatusMethodNotAllowed,
			method:     http.MethodPatch,
			url:        "/update/counter/test/1",
		},
		{
			name:       "correct request returns `http.StatusOK`",
			statusCode: http.StatusOK,
			method:     http.MethodPost,
			url:        "/update/counter/test/1",
		},
		{
			name:       "correct request returns `http.StatusOK`",
			statusCode: http.StatusOK,
			method:     http.MethodPost,
			url:        "/update/gauge/test/3.14",
		},
		{
			name: "request without metric name returns `http.StatusNotFound`",
			statusCode: http.StatusNotFound,
			method: http.MethodPost,
			url: "/update/counter/1",
		},
		{
			name: "request with invalid metric type returns `http.StatusBadRequest`",
			statusCode: http.StatusBadRequest,
			method: http.MethodPost,
			url: "/update/amogus/test/1",
		},
		{
			name: "request with invalid metric value (counter) return `http.StatusBadRequest`",
			statusCode: http.StatusBadRequest,
			method: http.MethodPost,
			url: "/update/counter/test/3.14",
		},
		{
			name: "request with invalid metric value (gauge) return `http.StatusBadRequest`",
			statusCode: http.StatusBadRequest,
			method: http.MethodPost,
			url: "/update/gauge/test/pi",
		},
	}

	for _, test := range tests {
		gauge := make(map[string]float64)
		counter := make(map[string]int64)
		storage := storage.NewMemStorage(gauge, counter)

		t.Run(test.name, func (t *testing.T) {
			request := httptest.NewRequest(test.method, test.url, nil)
			w := httptest.NewRecorder()
			UpdateMetrics(&storage)(w, request)

			res := w.Result()

			err := res.Body.Close()
			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != test.statusCode {
				t.Errorf("expected status code %v, but got %v", test.statusCode, res.StatusCode)
			}
		})
	}
}
