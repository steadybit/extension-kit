/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"compress/gzip"
	"github.com/klauspost/compress/gzhttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRequestTimeoutHeaderAware(t *testing.T) {
	tests := []struct {
		name                 string
		requestTimeoutHeader string
		wantedStatusCode     int
		wantsDeadline        bool
	}{
		{
			name:                 "Should apply no timeout if no header is set",
			requestTimeoutHeader: "",
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply no timeout if invalid header is set",
			requestTimeoutHeader: "foobar",
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply timeout if valid header is set",
			requestTimeoutHeader: "0.5",
			wantedStatusCode:     503,
			wantsDeadline:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			if len(tt.requestTimeoutHeader) > 0 {
				req.Header.Set("Request-Timeout", tt.requestTimeoutHeader)
			}

			rr := httptest.NewRecorder()
			RequestTimeoutHeaderAware(handler(t, tt.wantsDeadline)).ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantedStatusCode)
			}
		})
	}
}

func handler(t *testing.T, wantsDeadline bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(200)
		if _, ok := r.Context().Deadline(); wantsDeadline != ok {
			t.Errorf("Expected request context to have a deadline, but it didn't")
		}
	}
}

func TestIfNoneMatchHandler(t *testing.T) {
	tests := []struct {
		name                     string
		IfNoneMatchRequestHeader string
		ETag                     string
		wantedEtagResponseHeader string
		wantedStatusCode         int
	}{
		{
			name:                     "should return 200 and no header",
			IfNoneMatchRequestHeader: "",
			ETag:                     "",
			wantedEtagResponseHeader: "",
			wantedStatusCode:         200,
		},
		{
			name:                     "should return 200 and etag header",
			IfNoneMatchRequestHeader: "",
			ETag:                     "abcdef",
			wantedEtagResponseHeader: "abcdef",
			wantedStatusCode:         200,
		},
		{
			name:                     "should return 304 and etag header",
			IfNoneMatchRequestHeader: "abcdef",
			ETag:                     "abcdef",
			wantedEtagResponseHeader: "",
			wantedStatusCode:         304,
		},
		{
			name:                     "should return 200 and etag header",
			IfNoneMatchRequestHeader: "abcdef",
			ETag:                     "ghijkl",
			wantedEtagResponseHeader: "ghijkl",
			wantedStatusCode:         200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			if len(tt.IfNoneMatchRequestHeader) > 0 {
				req.Header.Set("If-None-Match", tt.IfNoneMatchRequestHeader)
			}

			rr := httptest.NewRecorder()
			IfNoneMatchHandler(func() string {
				return tt.ETag
			}, func(w http.ResponseWriter, r *http.Request, body []byte) {
				w.WriteHeader(200)
			})(rr, req, nil)

			if status := rr.Code; status != tt.wantedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantedStatusCode)
			}

			if etag := rr.Header().Get("ETag"); etag != tt.wantedEtagResponseHeader {
				t.Errorf("handler returned wrong etag header: got %v want %v", etag, tt.wantedEtagResponseHeader)
			}
		})
	}
}

func TestPanicRecovery(t *testing.T) {
	tests := []struct {
		name             string
		panic            bool
		wantedStatusCode int
	}{
		{
			name:             "should return 200",
			panic:            false,
			wantedStatusCode: 200,
		},
		{
			name:             "should return 500",
			panic:            true,
			wantedStatusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			PanicRecovery(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.panic {
					panic("test")
				}
				w.WriteHeader(200)
			})).ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantedStatusCode)
			}

		})
	}
}

func TestWriteBody(t *testing.T) {
	tests := []struct {
		name         string
		responseBody any
		wantedStatus int
		wantedBody   string
	}{
		{
			name: "should write body",
			responseBody: map[string]string{
				"key": "value",
			},
			wantedStatus: 200,
			wantedBody:   "{\"key\":\"value\"}\n",
		},
		{
			name:         "should fail writing body",
			responseBody: make(chan int),
			wantedStatus: 500,
			wantedBody:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			WriteBody(rr, tt.responseBody)

			assert.Equal(t, tt.wantedBody, rr.Body.String())
			assert.Equal(t, tt.wantedStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}

func TestGzipHandler(t *testing.T) {
	largeBody := `{"data":"` + strings.Repeat("x", 1500) + `"}`
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(largeBody))
	})

	t.Run("should compress response when Accept-Encoding gzip is set", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Accept-Encoding", "gzip")
		rr := httptest.NewRecorder()

		gzhttp.GzipHandler(inner).ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, "gzip", rr.Header().Get("Content-Encoding"))

		gr, err := gzip.NewReader(rr.Body)
		require.NoError(t, err)
		defer gr.Close()
		body, err := io.ReadAll(gr)
		require.NoError(t, err)
		assert.Equal(t, largeBody, string(body))
	})

	t.Run("should not compress response when Accept-Encoding gzip is not set", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		gzhttp.GzipHandler(inner).ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
		assert.Empty(t, rr.Header().Get("Content-Encoding"))
		assert.Equal(t, largeBody, rr.Body.String())
	})
}
